package suite

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/skudasov/linke2e/suite/contracts_client"
	"github.com/skudasov/linke2e/suite/node_client"
)

type ChainLinkSuite struct {
	Contracts  *contracts_client.ContractsInteractor
	NodeClient *node_client.NodeClient
}

func NewChainLinkSuite(cfg *SuiteConfig) *ChainLinkSuite {
	s := &ChainLinkSuite{
		Contracts:  contracts_client.NewContracts(cfg.ContractsConfig),
		NodeClient: node_client.NewNodeClient(cfg.NodeClientConfig),
	}
	s.Prepare()
	return s
}

// Prepare sets all required conditions, fund Chainlink node with eth/link, authorize in node API
func (m *ChainLinkSuite) Prepare() {
	m.NodeClient.Authorize()
	if m.Contracts.Cfg.NetworkType == contracts_client.GethNetwork {
		m.Contracts.GenerateAddresses(10)
		m.Contracts.DeployContracts()
		m.Contracts.FundConsumerWithLink()
	} else {
		m.Contracts.HardhatDeployerData()
	}
	m.Contracts.FundNodeWithEth(m.NodeClient.GetNodeEthAddr())
}

func (m *ChainLinkSuite) InteractionFromFiles(jobPath string) string {
	jobFile, err := os.Open(jobPath)
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(jobFile)

	var r map[string]interface{}
	if err := json.Unmarshal(byteValue, &r); err != nil {
		log.Fatal(err)
	}
	return m.CreateJobSpec(r)
}

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
