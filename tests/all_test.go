package tests

import (
	"testing"
	"time"

	"github.com/skudasov/linke2e/suite"
	"github.com/skudasov/linke2e/suite/contracts_client"
	"github.com/skudasov/linke2e/suite/mock_api"
	"github.com/skudasov/linke2e/suite/node_client"
)

func TestJobInteractions(t *testing.T) {
	h := suite.NewChainLinkSuite(&suite.Config{
		MockClientConfig: &mock_api.Config{
			Url: "http://localhost:9092",
		},
		NodeClientConfig: &node_client.Config{
			Url:      "http://localhost:6688",
			Login:    "notreal@fakeemail.ch",
			Password: "twochains",
		},
		ContractsConfig: &contracts_client.Config{
			NetworkType:          contracts_client.GethNetwork,
			RPCUrl:               "ws://localhost:8545",
			NetworkID:            "1337",
			HardhatPrivateKeyHex: "88b1ae324c09b52d8710bbe4c4f0b63f591e5eea508176a97d7f76b154f979b1",
			GethRootPassword:     "123456",
			GethRootAccountFile:  "../deploy/ethconfig/keystore/UTC--2021-04-01T13-44-29.143912000Z--01bca4c976169d4d8cceb227c08bab9e3775d7f5",
		},
		AssertConfig: &suite.AssertConfig{
			Attempts: 10,
			Delay:    1 * time.Second,
		},
	})

	var jobTests = []struct {
		name     string
		specFile string
		stubFile string
	}{
		// {"simple web", "data/simple_web.json", "data/simple_web_stub.json"},
		{"simple get", "data/simple_get.json", "data/simple_get_stub.json"},
	}

	for _, tt := range jobTests {
		t.Run(tt.name, func(t *testing.T) {
			defer h.ResetMock()
			stubMap := h.CreateStub(tt.stubFile)
			specMap := h.CreateSpec(tt.specFile)

			if specMap.Get("initiators.0.type").String() == "runlog" {
				h.Contracts.APIConsumerRequest(
					specMap.Get("jobID").String(),
					1e18,
					"http://host.docker.internal:9092/api_stub",
					"data",
					1,
				)
			} else {
				h.NodeClient.TriggerJobRun(specMap.Get("jobID").String())
			}
			h.AwaitAPICall(t, stubMap)
			h.AwaitDataOnChain(t, specMap, stubMap)
		})
	}
}
