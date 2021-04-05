package contracts_client

const (
	HardhatNetwork int = iota
	GethNetwork
)

type Config struct {
	RPCUrl               string
	NetworkID            string
	NetworkType          int
	HardhatPrivateKeyHex string
	GethRootAccountFile  string
	GethRootPassword     string
}
