package tests

import (
	"testing"

	"github.com/skudasov/linke2e/suite"
	"github.com/skudasov/linke2e/suite/contracts_client"
	"github.com/skudasov/linke2e/suite/node_client"
)

func TestJobInteractions(t *testing.T) {
	h := suite.NewChainLinkSuite(&suite.SuiteConfig{
		NodeClientConfig: &node_client.NodeClientConfig{
			Url:      "http://localhost:6688",
			Login:    "notreal@fakeemail.ch",
			Password: "twochains",
		},
		ContractsConfig: &contracts_client.ContractsConfig{
			NetworkType:          contracts_client.GethNetwork,
			RPCUrl:               "ws://localhost:8545",
			NetworkID:            "31337",
			HardhatPrivateKeyHex: "88b1ae324c09b52d8710bbe4c4f0b63f591e5eea508176a97d7f76b154f979b1",
			GethRootPassword:     "123456",
			GethRootAccountFile:  "../deploy/ethconfig/keystore/UTC--2021-04-01T13-44-29.143912000Z--01bca4c976169d4d8cceb227c08bab9e3775d7f5",
		},
	})

	var jobTests = []struct {
		name  string
		spec  string
		url   string
		times int
		out   string
	}{
		{"simple gen", "data/simple_get.json", "https://bitstamp.net/api/ticker/", 10, ""},
	}

	for _, tt := range jobTests {
		t.Run(tt.name, func(t *testing.T) {
			jobID := h.InteractionFromFiles(tt.spec)
			h.Contracts.APIConsumerRequest(jobID, tt.url, tt.times)
		})
	}
}
