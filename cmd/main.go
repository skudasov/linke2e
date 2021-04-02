package main

import (
	"time"

	"github.com/skudasov/linke2e/suite"
	"github.com/skudasov/linke2e/suite/contracts_client"
	"github.com/skudasov/linke2e/suite/node_client"
)

func main() {
	h := suite.NewChainLinkSuite(&suite.SuiteConfig{
		NodeClientConfig: &node_client.NodeClientConfig{
			Url:      "http://localhost:6688",
			Login:    "notreal@fakeemail.ch",
			Password: "twochains",
		},
		ContractsConfig: &contracts_client.ContractsConfig{
			NetworkType: contracts_client.GethNetwork,
			RPCUrl:      "ws://localhost:8545",
			NetworkID:   "31337",
			// hardhat
			HardhatManifestFile:  "contracts/manifest.json",
			HardhatPrivateKeyHex: "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
			// geth
			GethRootPassword:    "123456",
			GethRootAccountFile: "deploy/ethconfig/keystore/UTC--2021-04-01T13-44-29.143912000Z--01bca4c976169d4d8cceb227c08bab9e3775d7f5",
		},
	})
	jobID := h.CreateJobSpec(node_client.CreateJobBody{
		Initiators: []node_client.Initiator{
			{
				Type: "runlog",
				Params: map[string]string{
					"address": h.Contracts.GethDeployedContracts.MockOracleAddress.Hex(),
					// "address": h.Contracts.HardhatDeployerData.MockOracleAddress.String(),
				},
			},
		},
		Tasks: []node_client.Task{
			// {
			// 	Type:          "HTTPGetWithUnrestrictedNetworkAccess",
			// 	Confirmations: 0,
			// 	Params:        map[string]string{"get": "https://bitstamp.net/api/ticker/"},
			// },
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
		// Startat:    time.Now().Format(time.RFC3339),
		Endat:      nil,
		Minpayment: "1",
	})
	time.Sleep(1 * time.Second)
	// go h.LogListener()
	h.APIConsumerTest(jobID, "https://bitstamp.net/api/ticker/", 10)
	time.Sleep(5 * time.Second)
}
