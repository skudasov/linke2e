package mock_api

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	Cfg    *Config
	Client *resty.Client
}

func NewClient(cfg *Config) *Client {
	c := resty.New()
	c.SetHostURL(cfg.Url)
	if cfg.Debug {
		c.SetDebug(true)
	}
	return &Client{
		Cfg:    cfg,
		Client: c,
	}
}

func (m *Client) CheckCalledTimes(path string) int {
	var resCalled CallResponse
	_, err := m.Client.R().
		SetResult(&resCalled).
		Get(fmt.Sprintf("/check_call/%s", path))
	if err != nil {
		log.Fatal(err)
	}
	return resCalled.Times
}

func (m *Client) SetStubResponse(body map[string]interface{}) {
	_, err := m.Client.R().
		SetBody(body).
		Post("/set_stub_response")
	if err != nil {
		log.Fatal(err)
	}
}
