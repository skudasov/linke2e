package contracts_client

import (
	"context"
	"crypto/ecdsa"
	"io/ioutil"
	"log"
	"math/big"
	"testing"

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
func (m *ContractsInteractor) ShowAddresses(t *testing.T, amount int) {
	key := m.RootAccountFromFile()
	for i := 0; i < amount; i++ {
		t.Logf("address #%d: %s\n", i, crypto.CreateAddress(key.Address, uint64(i)).String())
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

func (m *ContractsInteractor) DeployContracts(t *testing.T) {
	k := m.RootAccountFromFile()
	m.PrintRootBalace(t, k)
	linkAddr, linkTx, mockLinkInstance, err := MockLink.DeployMockLink(
		m.DeployerTransactor(k.Address, k.PrivateKey),
		m.EthClient,
	)
	if err != nil {
		t.Error(err)
	}
	t.Logf("link deployed")
	t.Logf(linkAddr.Hex())
	t.Logf(linkTx.Hash().Hex())

	oracleAddr, oracleTx, mockOracleInstance, err := MockOracle.DeployMockOracle(
		m.DeployerTransactor(k.Address, k.PrivateKey),
		m.EthClient,
		linkAddr,
	)
	if err != nil {
		t.Error(err)
	}
	t.Logf("mock oracle deployed")
	t.Logf(oracleAddr.Hex())
	t.Logf(oracleTx.Hash().Hex())

	apiAddr, apiTx, apiConsumerInstance, err := APIConsumer.DeployAPIConsumer(
		m.DeployerTransactor(k.Address, k.PrivateKey),
		m.EthClient,
		linkAddr,
	)
	if err != nil {
		t.Error(err)
	}
	t.Logf("api consumer deployed")
	t.Logf(apiAddr.Hex())
	t.Logf(apiTx.Hash().Hex())
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
func (m *ContractsInteractor) FundConsumerWithLink(t *testing.T, link int64) {
	t.Logf("funding consumer with link")
	k := m.RootAccountFromFile()
	tx := m.DeployerTransactor(k.Address, k.PrivateKey)
	_, err := m.DeployedData.MockLink.Transfer(tx, m.DeployedData.APIConsumerAddress, big.NewInt(link))
	if err != nil {
		t.Error(err)
	}
}

// FundNodeWithEth fund node with eth
func (m *ContractsInteractor) FundNodeWithEth(t *testing.T, nodeAddr string) {
	nonce, err := m.EthClient.PendingNonceAt(context.Background(), m.DeployedData.RootAddress)
	if err != nil {
		t.Error(err)
	}
	gasPrice, err := m.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		t.Error(err)
	}
	chainID, err := m.EthClient.NetworkID(context.Background())
	if err != nil {
		t.Error(err)
	}
	var data []byte
	tx := types.NewTransaction(nonce, common.HexToAddress(nodeAddr), big.NewInt(1000000000000000000), uint64(300000), gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), m.DeployedData.RootPrivateKey)
	if err != nil {
		t.Error(err)
	}
	err = m.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		t.Error(err)
	}
	t.Logf("fund chainlink node tx hash: %s", signedTx.Hash().Hex())
}

func (m *ContractsInteractor) APIConsumerRequest(t *testing.T, jobID string, payment int64, url string, path string, times int) {
	var jobIDToSend [32]byte
	copy(jobIDToSend[:], jobID)
	t.Logf("job id hex: %s", hexutil.Encode([]byte(jobID)))

	t.Logf("calling APIConsumer.createRequestTo with oracle addr: %s", m.DeployedData.MockOracleAddress.Hex())
	res, err := m.DeployedData.APIConsumer.CreateRequestTo(
		m.DeployerTransactor(m.DeployedData.RootAddress, m.DeployedData.RootPrivateKey),
		m.DeployedData.MockOracleAddress,
		jobIDToSend,
		big.NewInt(payment),
		url,
		path,
		big.NewInt(int64(times)),
	)
	if err != nil {
		t.Error(err)
	}
	jsonPayload, err := res.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("receipt: %s", jsonPayload)
}

func (m *ContractsInteractor) SetFulfullmentPermission(t *testing.T, nodeAddr string) {
	if _, err := m.DeployedData.MockOracle.SetFulfillmentPermission(
		m.DeployerTransactor(m.DeployedData.RootAddress, m.DeployedData.RootPrivateKey),
		common.HexToAddress(nodeAddr),
		true,
	); err != nil {
		t.Error(err)
	}
}

func (m *ContractsInteractor) CheckAPIConsumerData() int64 {
	data, err := m.DeployedData.APIConsumer.Data(nil)
	if err != nil {
		log.Fatal(err)
	}
	return data.Int64()
}
