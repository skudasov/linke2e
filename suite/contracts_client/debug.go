package contracts_client

import (
	"context"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func (m *ContractsInteractor) PrintRootBalace(t *testing.T, k *keystore.Key) {
	bn, err := m.EthClient.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	balance, err := m.EthClient.BalanceAt(context.Background(), k.Address, big.NewInt(int64(bn)))
	if err != nil {
		log.Fatal(err)
	}
	t.Logf("balance of root: %d, addr: %s", balance, k.Address.String())
}
