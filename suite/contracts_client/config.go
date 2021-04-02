package contracts_client

const (
	HardhatNetwork int = iota
	GethNetwork
)

type ContractsConfig struct {
	RPCUrl      string
	NetworkID   string
	NetworkType int

	// hardhat
	HardhatManifestFile  string
	HardhatPrivateKeyHex string

	// geth
	GethRootAccountFile string
	GethRootPassword    string
}
