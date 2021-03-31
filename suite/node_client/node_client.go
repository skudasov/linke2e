package node_client

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type NodeClient struct {
	Cfg    *NodeClientConfig
	Client *resty.Client
}

func NewNodeClient(cfg *NodeClientConfig) *NodeClient {
	c := resty.New()
	c.SetHostURL(cfg.Url)
	if cfg.Debug {
		c.SetDebug(true)
	}
	return &NodeClient{
		Cfg:    cfg,
		Client: c,
	}
}

func (m *NodeClient) Authorize() {
	s := &Session{Email: m.Cfg.Login, Password: m.Cfg.Password}
	res, err := m.Client.R().
		SetBody(s).
		Post("/sessions")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("session created: %s", res)
}