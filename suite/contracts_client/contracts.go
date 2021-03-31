package contracts_client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/skudasov/linke2e/abi_build/consumer"
)

type ContractsInteractor struct {
	RPCUrl          string
	NetworkID       string
	manifest        map[string]interface{}
	ApiConsumerAddr common.Address
	DeployerData    *DeployerData
	EthClient       *ethclient.Client
}

func NewContracts(cfg *ContractsConfig) *ContractsInteractor {
	c := &ContractsInteractor{
		RPCUrl: cfg.RPCUrl,
	}
	var err error
	c.EthClient, err = ethclient.Dial(c.RPCUrl)
	if err != nil {
		log.Fatal(err)
	}
	c.DeployerData = NewDeployerData(cfg.DeployerPrivateKeyHex)
	c.NetworkID = cfg.NetworkID
	c.manifest = c.GetManifest()
	c.ApiConsumerAddr = common.HexToAddress(c.manifest[cfg.NetworkID].(map[string]interface{})["localhost"].(map[string]interface{})["contracts"].(map[string]interface{})["APIConsumer"].(map[string]interface{})["address"].(string))
	return c
}

func (m *ContractsInteractor) APIConsumerTest() {
	instance, err := consumer.NewAPIConsumer(m.ApiConsumerAddr, m.EthClient)
	if err != nil {
		log.Fatal(err)
	}
	res, err := instance.Owner(nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("res: %s", res)
}

func (m *ContractsInteractor) GetManifest() map[string]interface{} {
	f, err := os.Open("/Users/sergejkudasov/go/src/github.com/skudasov/linke2e/contracts/manifest.json")
	if err != nil {
		log.Fatal(err)
	}
	manifest, _ := ioutil.ReadAll(f)
	var manifestMap map[string]interface{}
	if err := json.Unmarshal(manifest, &manifestMap); err != nil {
		log.Fatal(err)
	}
	return manifestMap
}
