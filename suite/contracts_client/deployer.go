package contracts_client

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type DeployerData struct {
	PrivateKeyHex    string
	PrivateKey       *ecdsa.PrivateKey
	PublicKey        *ecdsa.PublicKey
	PublicKeyAddress common.Address
}

func NewDeployerData(hexKey string) *DeployerData {
	// TODO: hardcoded deployer private key for 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266 accounts[0]
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	return &DeployerData{
		PrivateKey:       privateKey,
		PublicKey:        publicKeyECDSA,
		PublicKeyAddress: crypto.PubkeyToAddress(*publicKeyECDSA),
	}
}
