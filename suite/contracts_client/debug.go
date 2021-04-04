package contracts_client

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/skudasov/linke2e/contracts_go_build/build/mockoracle"
)

func (m *ContractsInteractor) LogListener() {
	var ch = make(chan types.Log)
	ctx := context.Background()

	sub, err := m.EthClient.SubscribeFilterLogs(ctx, ethereum.FilterQuery{}, ch)

	if err != nil {
		log.Fatal(err)
	}

	tokenAbi, err := abi.JSON(strings.NewReader(mockoracle.MockOracleABI))

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

			err = tokenAbi.UnpackIntoInterface(&transferEvent, "MockOracle", eventLog.Data)

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
