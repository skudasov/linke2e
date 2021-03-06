package suite

import (
	"time"

	"github.com/skudasov/linke2e/suite/contracts_client"
	"github.com/skudasov/linke2e/suite/mock_api"
	"github.com/skudasov/linke2e/suite/node_client"
)

type Config struct {
	NodeClientConfig *node_client.Config
	ContractsConfig  *contracts_client.Config
	MockClientConfig *mock_api.Config
	AssertConfig     *AssertConfig
}

type AssertConfig struct {
	Attempts uint
	Delay    time.Duration
}
