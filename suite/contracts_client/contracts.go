package contracts_client

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/skudasov/linke2e/abi_build/consumer"
	"github.com/skudasov/linke2e/contracts_go_build/build/apiconsumer"
	"github.com/skudasov/linke2e/contracts_go_build/build/mocklink"
	"github.com/skudasov/linke2e/contracts_go_build/build/mockoracle"
	"github.com/skudasov/linke2e/suite/deployer"
)

// GethDeployedContracts provides data and stubs of deployed contracts on geth
type GethDeployedContracts struct {
	MockLinkAddress    common.Address
	MockLink           *mocklink.MockLink
	MockOracleAddress  common.Address
	MockOracle         *mockoracle.MockOracle
	APIConsumerAddress common.Address
	APIConsumer        *apiconsumer.APIConsumer
}

// ContractsInteractor wraps interactions with contracts for both hardhat and geth
type ContractsInteractor struct {
	Cfg       *ContractsConfig
	manifest  map[string]interface{}
	EthClient *ethclient.Client

	// geth deps
	GethDeployedContracts *GethDeployedContracts

	// hardhat deps
	ApiConsumerAddr     common.Address
	HardhatDeployerData *deployer.HardhatDeployerData
}

// NewContracts setups client connection with blockchain network
func NewContracts(cfg *ContractsConfig) *ContractsInteractor {
	c := &ContractsInteractor{
		Cfg: cfg,
	}
	var err error
	c.EthClient, err = ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		log.Fatal(err)
	}
	if cfg.NetworkType == HardhatNetwork {
		c.HardhatDeployerData = deployer.NewHardhatDeployerData(cfg.HardhatPrivateKeyHex)
	}
	return c
}

// GethRootAccount loads geth root account info
func (m *ContractsInteractor) GethRootAccount() *keystore.Key {
	jsonBytes, err := ioutil.ReadFile(m.Cfg.GethRootAccountFile)
	if err != nil {
		log.Fatal(err)
	}
	k, err := keystore.DecryptKey(jsonBytes, m.Cfg.GethRootPassword)
	if err != nil {
		log.Fatal(err)
	}
	bn, err := m.EthClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	balance, err := m.EthClient.BalanceAt(context.Background(), k.Address, big.NewInt(int64(bn)))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("balance of root: %d, addr: %s", balance, k.Address.String())
	return k
}

// GenerateAddresses helper method to list generated addresses order
func (m *ContractsInteractor) GenerateAddresses(amount int) {
	key := m.GethRootAccount()
	for i := 0; i < 5; i++ {
		log.Printf("address #%d: %s\n", i, crypto.CreateAddress(key.Address, uint64(amount)).String())
	}
}

// DeployerTransactor creates default root account signer opts
func (m *ContractsInteractor) DeployerTransactor(rootAddr common.Address, rootPrivateKey *ecdsa.PrivateKey) *bind.TransactOpts {
	nonce, err := m.EthClient.PendingNonceAt(context.Background(), rootAddr)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := m.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx := bind.NewKeyedTransactor(rootPrivateKey)
	tx.Nonce = big.NewInt(int64(nonce))
	tx.Value = big.NewInt(0)
	tx.GasLimit = uint64(3000000)
	tx.GasPrice = gasPrice
	return tx
}

func (m *ContractsInteractor) DeployContracts() {
	k := m.GethRootAccount()
	linkAddr, linkTx, mockLinkInstance, err := mocklink.DeployMockLink(
		m.DeployerTransactor(k.Address, k.PrivateKey),
		m.EthClient,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("link deployed")
	log.Println(linkAddr.Hex())
	log.Println(linkTx.Hash().Hex())

	oracleAddr, oracleTx, mockOracleInstance, err := mockoracle.DeployMockOracle(
		m.DeployerTransactor(k.Address, k.PrivateKey),
		m.EthClient,
		linkAddr,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("mock oracle deployed")
	log.Println(oracleAddr.Hex())
	log.Println(oracleTx.Hash().Hex())

	apiAddr, apiTx, apiConsumerInstance, err := apiconsumer.DeployAPIConsumer(
		m.DeployerTransactor(k.Address, k.PrivateKey),
		m.EthClient,
		linkAddr,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("api consumer deployed")
	log.Println(apiAddr.Hex())
	log.Println(apiTx.Hash().Hex())
	m.GethDeployedContracts = &GethDeployedContracts{
		MockLinkAddress:    linkAddr,
		MockLink:           mockLinkInstance,
		MockOracleAddress:  oracleAddr,
		MockOracle:         mockOracleInstance,
		APIConsumerAddress: apiAddr,
		APIConsumer:        apiConsumerInstance,
	}
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
	f, err := os.Open(m.Cfg.HardhatManifestFile)
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
