package contracts_client

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/skudasov/linke2e/abi_build/consumer"
)

func (m *ContractsInteractor) APIConsumerRequest(jobID string, url string, times int) {
	transactor := m.DeployerTransactor(m.DeployedData.RootAddress, m.DeployedData.RootPrivateKey)
	instance, err := consumer.NewAPIConsumer(m.DeployedData.APIConsumerAddress, m.EthClient)
	if err != nil {
		log.Fatal(err)
	}

	var jobIDToSend [32]byte
	log.Printf("job id hex: %s", hexutil.Encode([]byte(jobID)))
	copy(jobIDToSend[:], jobID)

	log.Printf("creating tx to mock oracle: %s", m.DeployedData.MockOracleAddress.Hex())
	res, err := instance.CreateRequestTo(
		transactor,
		m.DeployedData.MockOracleAddress,
		jobIDToSend,
		big.NewInt(10000000000000000),
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
}
