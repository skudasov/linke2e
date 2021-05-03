package node_client

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
)

type NodeClient struct {
	Cfg    *Config
	Client *resty.Client
}

func NewNodeClient(cfg *Config) *NodeClient {
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

func (m *NodeClient) Authorize(t *testing.T) {
	s := &Session{Email: m.Cfg.Login, Password: m.Cfg.Password}
	res, err := m.Client.R().
		SetBody(s).
		Post("/sessions")
	if err != nil {
		t.Error(err)
	}
	t.Logf("session created: %s", res)
}

func (m *NodeClient) GetNodeEthAddr(t *testing.T) string {
	var respBody GetBalancesBodyResponse
	res, err := m.Client.R().
		SetResult(&respBody).
		Get("/v2/user/balances")
	if err != nil {
		t.Error(err)
	}
	ethAddr := respBody.Data[0].ID
	t.Logf("node balances response: %s", res)
	t.Logf("node eth id: %s", ethAddr)
	return ethAddr
}

func (m *NodeClient) TriggerJobRun(t *testing.T, jobID string) {
	var respBody map[string]interface{}
	res, err := m.Client.R().
		SetResult(&respBody).
		Post(fmt.Sprintf("/v2/specs/%s/runs", jobID))
	if err != nil {
		t.Error(err)
	}
	t.Logf("trigger response for job: %s: %v", jobID, res)
}
