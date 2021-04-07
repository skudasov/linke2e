package contracts_client

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
)

func (m *ContractsInteractor) PrintRootBalace(k *keystore.Key) {
	bn, err := m.EthClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	balance, err := m.EthClient.BalanceAt(context.Background(), k.Address, big.NewInt(int64(bn)))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("balance of root: %d, addr: %s", balance, k.Address.String())
}

func (m *ContractsInteractor) WatchConsumerRequests(jobID string) {
	var jobIDToSend [32]byte
	copy(jobIDToSend[:], jobID)
	var ji = make([][32]byte, 1)
	ji[0] = jobIDToSend
	iter, err := m.DeployedData.APIConsumer.FilterChainlinkRequested(nil, ji)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("chainlink requested: %v", iter.Event)
	for iter.Next() {
		log.Printf("chainlink requested: %v", iter.Event)
	}
}

func (m *ContractsInteractor) DumpVars() {
	cbAddr, _ := m.DeployedData.MockOracle.MemoCallbackAddr(nil)
	cbFn, _ := m.DeployedData.MockOracle.MemoCallbackFn(nil)
	selector, _ := m.DeployedData.APIConsumer.Selector(nil)
	log.Printf("cb addr: %s, cb fn: %s, selector provided: %s", cbAddr.Hex(), common.Bytes2Hex(cbFn[:]), common.Bytes2Hex(selector[:]))
}

func (m *ContractsInteractor) DumpEvents() {
	iter, _ := m.DeployedData.MockOracle.FilterOracleRequest(nil, [][32]byte{})
	for iter.Next() {
		log.Printf("OracleRequest: %v", iter.Event)
	}
	iter2, _ := m.DeployedData.MockOracle.FilterCancelOracleRequest(nil, nil)
	for iter.Next() {
		log.Printf("CancelOracleRequest: %v", iter2.Event)
	}
	iter3, _ := m.DeployedData.APIConsumer.FilterChainlinkFulfilled(nil, nil)
	for iter.Next() {
		log.Printf("ChainlinkFulfilled: %v", iter3.Event)
	}
}
