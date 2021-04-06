package contracts_client

import (
	"context"
	"crypto/ecdsa"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/skudasov/linke2e/contracts_go_build/build/APIConsumer"
	"github.com/skudasov/linke2e/contracts_go_build/build/MockLink"
	"github.com/skudasov/linke2e/contracts_go_build/build/MockOracle"
)

// DeployedData provides data and stubs of deployed contracts
type DeployedData struct {
	RootAddress        common.Address
	RootPrivateKey     *ecdsa.PrivateKey
	MockLinkAddress    common.Address
	MockLink           *MockLink.MockLink
	MockOracleAddress  common.Address
	MockOracle         *MockOracle.MockOracle
	APIConsumerAddress common.Address
	APIConsumer        *APIConsumer.APIConsumer
}

// ContractsInteractor wraps interactions with contracts for both hardhat and geth
type ContractsInteractor struct {
	Cfg          *Config
	EthClient    *ethclient.Client
	DeployedData *DeployedData
}

// NewContracts setups client connection with blockchain network
func NewContracts(cfg *Config) *ContractsInteractor {
	c := &ContractsInteractor{
		Cfg: cfg,
	}
	var err error
	c.EthClient, err = ethclient.Dial(cfg.RPCUrl)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

// RootAccountFromFile loads geth root account info
func (m *ContractsInteractor) RootAccountFromFile() *keystore.Key {
	jsonBytes, err := ioutil.ReadFile(m.Cfg.GethRootAccountFile)
	if err != nil {
		log.Fatal(err)
	}
	k, err := keystore.DecryptKey(jsonBytes, m.Cfg.GethRootPassword)
	if err != nil {
		log.Fatal(err)
	}
	return k
}

// ShowAddresses helper method to list generated addresses order
func (m *ContractsInteractor) ShowAddresses(amount int) {
	key := m.RootAccountFromFile()
	for i := 0; i < amount; i++ {
		log.Printf("address #%d: %s\n", i, crypto.CreateAddress(key.Address, uint64(i)).String())
	}
}

// DeployerTransactor creates default root account signer opts with a new nonce
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

func (m *ContractsInteractor) HardhatDeployerData() {
	privateKey, err := crypto.HexToECDSA(m.Cfg.HardhatPrivateKeyHex)
	if err != nil {
		log.Fatal(err)
	}
	m.DeployedData = &DeployedData{
		RootAddress:        common.HexToAddress("0x0757f2f2b672454d44a0a435aef0718f8528414f"),
		RootPrivateKey:     privateKey,
		MockLinkAddress:    common.HexToAddress("0xc5F64761A05aef6277bCCEf13B926df361ca56e8"),
		MockLink:           nil,
		MockOracleAddress:  common.HexToAddress("0x451B877E675e86b72da1525C2932fC5B4213f581"),
		MockOracle:         nil,
		APIConsumerAddress: common.HexToAddress("0xaC3dE07f7F44C0D722f1806d9159c494bAfa2145"),
		APIConsumer:        nil,
	}
}

func (m *ContractsInteractor) DeployContracts() {
	k := m.RootAccountFromFile()
	m.PrintRootBalace(k)
	linkAddr, linkTx, mockLinkInstance, err := MockLink.DeployMockLink(
		m.DeployerTransactor(k.Address, k.PrivateKey),
		m.EthClient,
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("link deployed")
	log.Println(linkAddr.Hex())
	log.Println(linkTx.Hash().Hex())

	oracleAddr, oracleTx, mockOracleInstance, err := MockOracle.DeployMockOracle(
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

	apiAddr, apiTx, apiConsumerInstance, err := APIConsumer.DeployAPIConsumer(
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
	m.DeployedData = &DeployedData{
		RootAddress:        k.Address,
		RootPrivateKey:     k.PrivateKey,
		MockLinkAddress:    linkAddr,
		MockLink:           mockLinkInstance,
		MockOracleAddress:  oracleAddr,
		MockOracle:         mockOracleInstance,
		APIConsumerAddress: apiAddr,
		APIConsumer:        apiConsumerInstance,
	}
}

// FundConsumerWithLink fund consumer contract with link, used only with geth deployment
func (m *ContractsInteractor) FundConsumerWithLink(link int64) {
	log.Printf("funding consumer with link")
	k := m.RootAccountFromFile()
	tx := m.DeployerTransactor(k.Address, k.PrivateKey)
	_, err := m.DeployedData.MockLink.Transfer(tx, m.DeployedData.APIConsumerAddress, big.NewInt(link))
	if err != nil {
		log.Fatal(err)
	}
}

// FundNodeWithEth fund node with eth
func (m *ContractsInteractor) FundNodeWithEth(nodeAddr string) {
	nonce, err := m.EthClient.PendingNonceAt(context.Background(), m.DeployedData.RootAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := m.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := m.EthClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var data []byte
	tx := types.NewTransaction(nonce, common.HexToAddress(nodeAddr), big.NewInt(1000000000000000000), uint64(300000), gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), m.DeployedData.RootPrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = m.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("fund chainlink node tx hash: %s", signedTx.Hash().Hex())
}

func (m *ContractsInteractor) APIConsumerRequest(jobID string, url string, path string, times int) {
	var jobIDToSend [32]byte
	copy(jobIDToSend[:], jobID)
	log.Printf("job id hex: %s", hexutil.Encode([]byte(jobID)))

	log.Printf("calling APIConsumer.createRequestTo with oracle addr: %s", m.DeployedData.MockOracleAddress.Hex())
	res, err := m.DeployedData.APIConsumer.CreateRequestTo(
		m.DeployerTransactor(m.DeployedData.RootAddress, m.DeployedData.RootPrivateKey),
		m.DeployedData.MockOracleAddress,
		jobIDToSend,
		big.NewInt(10000000000000000),
		url,
		path,
		big.NewInt(int64(times)),
	)
	if err != nil {
		log.Fatal(err)
	}
	jsonPayload, err := res.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("receipt: %s", jsonPayload)
}

func (m *ContractsInteractor) CheckAPIConsumerData() int64 {
	data, err := m.DeployedData.APIConsumer.Data(nil)
	if err != nil {
		log.Fatal(err)
	}
	return data.Int64()
}
