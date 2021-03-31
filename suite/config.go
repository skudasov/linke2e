package suite

import (
	"github.com/skudasov/linke2e/suite/contracts_client"
	"github.com/skudasov/linke2e/suite/node_client"
)

type SuiteConfig struct {
	NodeClientConfig *node_client.NodeClientConfig
	ContractsConfig  *contracts_client.ContractsConfig
}
