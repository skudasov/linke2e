package suite

import (
	"log"
	"testing"
	"time"

	"github.com/avast/retry-go"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"

	"github.com/skudasov/linke2e/suite/contracts_client"
	"github.com/skudasov/linke2e/suite/mock_api"
	"github.com/skudasov/linke2e/suite/node_client"
)

func init() {
	retry.DefaultAttempts = 10
	retry.DefaultDelayType = func(n uint, err error, config *retry.Config) time.Duration {
		return 1 * time.Second
	}
}

type ChainLinkSuite struct {
	Contracts  *contracts_client.ContractsInteractor
	NodeClient *node_client.NodeClient
	MockClient *mock_api.Client
}

func NewChainLinkSuite(cfg *Config) *ChainLinkSuite {
	s := &ChainLinkSuite{
		Contracts:  contracts_client.NewContracts(cfg.ContractsConfig),
		NodeClient: node_client.NewNodeClient(cfg.NodeClientConfig),
		MockClient: mock_api.NewClient(cfg.MockClientConfig),
	}
	s.Prepare()
	return s
}

// Prepare sets all required conditions, fund Chainlink node with eth/link, authorize in node API
func (m *ChainLinkSuite) Prepare() {
	m.NodeClient.Authorize()
	go mock_api.Start()
	if m.Contracts.Cfg.NetworkType == contracts_client.GethNetwork {
		m.Contracts.ShowAddresses(10)
		m.Contracts.DeployContracts()
		m.Contracts.FundConsumerWithLink(2e18)
	} else {
		m.Contracts.HardhatDeployerData()
	}
	m.Contracts.FundNodeWithEth(m.NodeClient.GetNodeEthAddr())
}

// CreateSpec creates job from spec file
func (m *ChainLinkSuite) CreateSpec(jobPath string) gjson.Result {
	g := GJSONFromFile(jobPath)
	rawMap := GJSONToMap(g)
	jobID := m.CreateJobSpec(rawMap)
	rawMap["jobID"] = jobID
	return GJSONFromMap(rawMap)
}

// CreateStub sets stub response from file
func (m *ChainLinkSuite) CreateStub(stubPath string) gjson.Result {
	g := GJSONFromFile(stubPath)
	rawMap := GJSONToMap(g)
	m.MockClient.SetStubResponse(rawMap["response"].(map[string]interface{}))
	return GJSONFromMap(rawMap)
}

// AwaitAPICall awaits that API called N times
func (m *ChainLinkSuite) AwaitAPICall(t *testing.T, stubMap gjson.Result) {
	if err := retry.Do(func() error {
		if int(stubMap.Get("times").Int()) != m.MockClient.CheckCalledTimes("api_stub") {
			log.Printf("checking stub was called")
			return errors.New("retrying awaiting for api_stub calls")
		}
		log.Printf("stub called")
		return nil
	}); err != nil {
		log.Printf("stub expectations failed")
		t.Fail()
	}
}

// AwaitDataOnChain awaits data on chain delivered
func (m *ChainLinkSuite) AwaitDataOnChain(t *testing.T, jobMap gjson.Result, stubMap gjson.Result) {
	data := stubMap.Get("response.data").Float()
	if err := retry.Do(func() error {
		d := m.Contracts.CheckAPIConsumerData()
		if d != int64(data) {
			log.Printf("awaiting for data: %d", d)
			return errors.New("retrying awaiting for data in contract")
		}
		return nil
	}); err != nil {
		log.Printf("on-chain data expectations failed")
		t.Fail()
	}
}

// CreateJobSpec creates job spec from map
func (m *ChainLinkSuite) CreateJobSpec(body map[string]interface{}) string {
	var respBody node_client.CreateJobBodyResponse
	res, err := m.NodeClient.Client.R().
		SetBody(body).
		SetResult(&respBody).
		Post("/v2/specs")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("job response: %s", res)
	log.Printf("job %s created", respBody.Data.ID)
	return respBody.Data.ID
}

// ResetMock resets mock response and call times
func (m *ChainLinkSuite) ResetMock() {
	log.Printf("resetting mock")
	m.MockClient.Reset()
}
