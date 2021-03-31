package suite

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/skudasov/linke2e/abi_build/consumer"
	"github.com/skudasov/linke2e/suite/contracts_client"
	"github.com/skudasov/linke2e/suite/node_client"
)

type ChainLinkSuite struct {
	Contracts  *contracts_client.ContractsInteractor
	NodeClient *node_client.NodeClient
}

func NewChainLinkSuite(cfg *SuiteConfig) *ChainLinkSuite {
	s := &ChainLinkSuite{
		Contracts:  contracts_client.NewContracts(cfg.ContractsConfig),
		NodeClient: node_client.NewNodeClient(cfg.NodeClientConfig),
	}
	s.Prepare()
	return s
}

// Prepare sets all required conditions, fund Chainlink node with eth/link, authorize in node API
func (m *ChainLinkSuite) Prepare() {
	m.NodeClient.Authorize()
	m.FundNodeWithEth()
}

func (m *ChainLinkSuite) CreateJobSpec(body node_client.CreateJobBody) string {
	var respBody node_client.CreateJobBodyResponse
	res, err := m.NodeClient.Client.R().
		SetBody(body).
		SetResult(&respBody).
		Post("/v2/specs")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("job response: %s", res)
	log.Printf("job %s created", respBody.Data.ID)
	return respBody.Data.ID
}

func (m *ChainLinkSuite) GetNodeEthAddr() string {
	var respBody node_client.GetBalancesBodyResponse
	res, err := m.NodeClient.Client.R().
		SetResult(&respBody).
		Get("/v2/user/balances")
	if err != nil {
		log.Fatal(err)
	}
	ethAddr := respBody.Data[0].ID
	log.Printf("node balances response: %s", res)
	log.Printf("node eth id: %s", ethAddr)
	return ethAddr
}

func (m *ChainLinkSuite) LogListener(contractAddr string) {
	contractAddress := common.HexToAddress(contractAddr)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	var ch = make(chan types.Log)
	ctx := context.Background()

	sub, err := m.Contracts.EthClient.SubscribeFilterLogs(ctx, query, ch)

	if err != nil {
		log.Fatal(err)
	}

	tokenAbi, err := abi.JSON(strings.NewReader(consumer.APIConsumerABI))

	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case eventLog := <-ch:
			var transferEvent struct {
				From  common.Address
				To    common.Address
				Value *big.Int
			}

			err = tokenAbi.UnpackIntoInterface(&transferEvent, "Transfer", eventLog.Data)

			if err != nil {
				log.Println("Failed to unpack")
				continue
			}

			transferEvent.From = common.BytesToAddress(eventLog.Topics[1].Bytes())
			transferEvent.To = common.BytesToAddress(eventLog.Topics[2].Bytes())

			log.Println("From", transferEvent.From.Hex())
			log.Println("To", transferEvent.To.Hex())
			log.Println("Value", transferEvent.Value)
		}
	}
}

func (m *ChainLinkSuite) FundNodeWithEth() {
	ethAddr := m.GetNodeEthAddr()
	nonce, err := m.Contracts.EthClient.PendingNonceAt(context.Background(), m.Contracts.DeployerData.PublicKeyAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := m.Contracts.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainID, err := m.Contracts.EthClient.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var data []byte
	tx := types.NewTransaction(nonce, common.HexToAddress(ethAddr), big.NewInt(1000000000000000000), uint64(300000), gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), m.Contracts.DeployerData.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Contracts.EthClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("fund chainlink node tx hash: %s", signedTx.Hash().Hex())
}

func (m *ChainLinkSuite) GetLastBlock() uint64 {
	blockNum, err := m.Contracts.EthClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return blockNum
}

func (m *ChainLinkSuite) APIConsumerTest(oracleAddr string, jobID string, url string, times int) {
	nonce, err := m.Contracts.EthClient.PendingNonceAt(context.Background(), m.Contracts.DeployerData.PublicKeyAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := m.Contracts.EthClient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	instance, err := consumer.NewAPIConsumer(m.Contracts.ApiConsumerAddr, m.Contracts.EthClient)
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(m.Contracts.DeployerData.PrivateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(300000)

	var jobIDToSend [32]byte
	log.Printf("job id hex: %s", hexutil.Encode([]byte(jobID)))
	copy(jobIDToSend[:], jobID)

	res, err := instance.CreateRequestTo(
		auth,
		common.HexToAddress(oracleAddr),
		jobIDToSend,
		big.NewInt(1000000000000000000),
		url,
		"",
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
	receipt, err := m.Contracts.EthClient.TransactionReceipt(context.Background(), res.Hash())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("txHash: %s", receipt.TxHash)
}
