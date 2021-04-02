package deployer

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type HardhatDeployerData struct {
	PrivateKeyHex      string
	PrivateKey         *ecdsa.PrivateKey
	PublicKey          *ecdsa.PublicKey
	PublicKeyAddress   common.Address
	MockLinkAddress    common.Address
	MockOracleAddress  common.Address
	APIConsumerAddress common.Address
}

func NewHardhatDeployerData(hexKey string) *HardhatDeployerData {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	return &HardhatDeployerData{
		PrivateKeyHex:      hexKey,
		PrivateKey:         privateKey,
		PublicKey:          publicKeyECDSA,
		PublicKeyAddress:   crypto.PubkeyToAddress(*publicKeyECDSA),
		MockLinkAddress:    common.HexToAddress("0x5fbdb2315678afecb367f032d93f642f64180aa3"),
		MockOracleAddress:  common.HexToAddress("0xe7f1725e7734ce288f8367e1bb143e90bb3f0512"),
		APIConsumerAddress: common.HexToAddress("0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0"),
	}
}
