package contracts_client

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/skudasov/linke2e/contracts_go_build/build/apiconsumer"
)

func (m *ContractsInteractor) APIConsumerRequest(jobID string, url string, path string, times int) {
	transactor := m.DeployerTransactor(m.DeployedData.RootAddress, m.DeployedData.RootPrivateKey)
	instance, err := apiconsumer.NewAPIConsumer(m.DeployedData.APIConsumerAddress, m.EthClient)
	if err != nil {
		log.Fatal(err)
	}

	var jobIDToSend [32]byte
	log.Printf("job id hex: %s", hexutil.Encode([]byte(jobID)))
	copy(jobIDToSend[:], jobID)

	log.Printf("calling APIConsumer.createRequestTo with oracle addr: %s", m.DeployedData.MockOracleAddress.Hex())
	res, err := instance.CreateRequestTo(
		transactor,
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
