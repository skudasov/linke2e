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
		PrivateKeyHex:    hexKey,
		PrivateKey:       privateKey,
		PublicKey:        publicKeyECDSA,
		PublicKeyAddress: crypto.PubkeyToAddress(*publicKeyECDSA),
		// old addresses
		// MockLinkAddress:    common.HexToAddress("0x5fbdb2315678afecb367f032d93f642f64180aa3"),
		MockLinkAddress: common.HexToAddress("0xc5F64761A05aef6277bCCEf13B926df361ca56e8"),
		// MockOracleAddress:  common.HexToAddress("0xe7f1725e7734ce288f8367e1bb143e90bb3f0512"),
		MockOracleAddress: common.HexToAddress("0x451B877E675e86b72da1525C2932fC5B4213f581"),
		// APIConsumerAddress: common.HexToAddress("0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0"),
		APIConsumerAddress: common.HexToAddress("0xaC3dE07f7F44C0D722f1806d9159c494bAfa2145"),
	}
}
