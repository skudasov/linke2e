package contracts_client

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
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

func (m *ContractsInteractor) WatchOracleEvents(jobID string) {
	var jobIDToSend [32]byte
	copy(jobIDToSend[:], jobID)
	var ji = make([][32]byte, 1)
	ji[0] = jobIDToSend
	iter, err := m.DeployedData.MockOracle.FilterOracleRequest(nil, ji)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("oracle requested: %v", iter.Event)
	for iter.Next() {
		log.Printf("oracle requested: %v", iter.Event)
	}
}
