package gdoc

import (
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	APIpath string       `json:"gdoc-api_path"`
	Token   string       `json:"gdoc-api_key"`
	client  *http.Client `json:"-"`
}

const timeout = time.Second * 10

func NewClient(conf []byte) (*Client, error) {
	var cli Client

	if err := json.Unmarshal(conf, &cli); err != nil {
		return nil, err
	}
	cli.client = &http.Client{Timeout: timeout}
	return &cli, nil
}
