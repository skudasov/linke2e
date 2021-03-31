package main

import (
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/skudasov/linke2e/suite"
	"github.com/skudasov/linke2e/suite/contracts_client"
	"github.com/skudasov/linke2e/suite/node_client"
)

func main() {
	mockOracleAddr := "0xe7f1725e7734ce288f8367e1bb143e90bb3f0512"
	h := suite.NewChainLinkSuite(&suite.SuiteConfig{
		NodeClientConfig: &node_client.NodeClientConfig{
			Url:      "http://localhost:6688",
			Login:    "notreal@fakeemail.ch",
			Password: "twochains",
		},
		ContractsConfig: &contracts_client.ContractsConfig{
			RPCUrl:                "http://localhost:8545",
			NetworkID:             "31337",
			DeployerPrivateKeyHex: "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		},
	})
	jobID := h.CreateJobSpec(node_client.CreateJobBody{
		Initiators: []node_client.Initiator{
			{
				Type:   "runlog",
				Params: map[string]string{"address": common.HexToAddress(mockOracleAddr).String(), "fromBlock": strconv.FormatUint(h.GetLastBlock(), 10)},
			},
		},
		Tasks: []node_client.Task{
			{
				Type:          "HTTPGet",
				Confirmations: 0,
				Params:        map[string]string{"get": "https://bitstamp.net/api/ticker/"},
			},
			{
				Type:   "JSONParse",
				Params: map[string]string{"path": "last"},
			},
			{
				Type:   "Multiply",
				Params: map[string]string{"times": "100"},
			},
			{
				Type:   "EthUint256",
				Params: map[string]string{},
			},
			{
				Type:   "EthTx",
				Params: map[string]string{},
			},
		},
		Startat:    time.Now().Format(time.RFC3339),
		Endat:      nil,
		Minpayment: "1",
	})
	time.Sleep(5 * time.Second)
	h.APIConsumerTest(mockOracleAddr, jobID, "http://ya.ru", 10)
}
