package suite

import (
	"testing"
	"time"

	"github.com/avast/retry-go"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"

	"github.com/skudasov/linke2e/suite/contracts_client"
	"github.com/skudasov/linke2e/suite/mock_api"
	"github.com/skudasov/linke2e/suite/node_client"
)

type ChainLinkSuite struct {
	Contracts    *contracts_client.ContractsInteractor
	NodeClient   *node_client.NodeClient
	MockClient   *mock_api.Client
	AssertConfig *AssertConfig
}

func NewChainLinkSuite(cfg *Config) *ChainLinkSuite {
	s := &ChainLinkSuite{
		Contracts:    contracts_client.NewContracts(cfg.ContractsConfig),
		NodeClient:   node_client.NewNodeClient(cfg.NodeClientConfig),
		MockClient:   mock_api.NewClient(cfg.MockClientConfig),
		AssertConfig: cfg.AssertConfig,
	}
	retry.DefaultAttempts = cfg.AssertConfig.Attempts
	retry.DefaultDelayType = func(n uint, err error, config *retry.Config) time.Duration {
		return cfg.AssertConfig.Delay
	}
	return s
}

// Prepare sets all required conditions, fund Chainlink node with eth/link, authorize in node API
func (m *ChainLinkSuite) Prepare(t *testing.T) {
	go mock_api.Start()
	m.NodeClient.Authorize(t)
	m.Contracts.ShowAddresses(t, 10)
	m.Contracts.DeployContracts(t)
	m.Contracts.FundConsumerWithLink(t, 2e18)
	m.Contracts.FundNodeWithEth(t, m.NodeClient.GetNodeEthAddr(t))
	m.Contracts.SetFulfullmentPermission(t, m.NodeClient.GetNodeEthAddr(t))
}

// CreateSpec creates job from spec file
func (m *ChainLinkSuite) CreateSpec(t *testing.T, jobPath string) gjson.Result {
	g := GJSONFromFile(jobPath)
	rawMap := GJSONToMap(g)
	jobID := m.CreateJobSpec(t, rawMap)
	rawMap["jobID"] = jobID
	return GJSONFromMap(rawMap)
}

// CreateStub sets stub response from file
func (m *ChainLinkSuite) CreateStub(t *testing.T, stubPath string) gjson.Result {
	g := GJSONFromFile(stubPath)
	rawMap := GJSONToMap(g)
	m.MockClient.SetStubResponse(t, rawMap["response"].(map[string]interface{}))
	return GJSONFromMap(rawMap)
}

// AwaitAPICall awaits that API called N times
func (m *ChainLinkSuite) AwaitAPICall(t *testing.T, stubMap gjson.Result) {
	if err := retry.Do(func() error {
		if int(stubMap.Get("times").Int()) != m.MockClient.CheckCalledTimes(t, "api_stub") {
			t.Logf("checking stub was called")
			return errors.New("retrying awaiting for api_stub calls")
		}
		t.Logf("stub called")
		return nil
	}); err != nil {
		t.Error(errors.Wrap(err, "stub expectations failed"))
	}
}

// AwaitDataOnChain awaits data on chain delivered
func (m *ChainLinkSuite) AwaitDataOnChain(t *testing.T, _ gjson.Result, stubMap gjson.Result) {
	expectedData := stubMap.Get("response.data").Float()
	if err := retry.Do(func() error {
		actualData := m.Contracts.CheckAPIConsumerData()
		t.Logf("expecting data: %d, have: %d", int64(expectedData), actualData)
		if actualData != int64(expectedData) {
			return errors.New("retrying awaiting for data in contract")
		}
		t.Logf("on chain data expectation pass")
		return nil
	}); err != nil {
		t.Error(errors.Wrap(err, "on-chain data expectations failed"))
	}
}

// CreateJobSpec creates job spec from map
func (m *ChainLinkSuite) CreateJobSpec(t *testing.T, body map[string]interface{}) string {
	var respBody node_client.CreateJobBodyResponse
	res, err := m.NodeClient.Client.R().
		SetBody(body).
		SetResult(&respBody).
		Post("/v2/specs")
	if err != nil {
		t.Error(err)
	}
	t.Logf("job response: %s", res)
	t.Logf("job %s created", respBody.Data.ID)
	return respBody.Data.ID
}

// ResetMock resets mock response and call times
func (m *ChainLinkSuite) ResetMock(t *testing.T) {
	t.Logf("resetting mock")
	m.MockClient.Reset(t)
}
