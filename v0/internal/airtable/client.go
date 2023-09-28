package airtable

import (
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	APIpath string       `json:"airtable-api_path"`
	Token   string       `json:"airtable-api_key"`
	Client  *http.Client `json:"-"`
}

const timeout = time.Second * 10

func NewClient(data []byte) (*Client, error) {
	var cli Client

	if err := json.Unmarshal(data, &cli); err != nil {
		return nil, err
	}
	cli.Client = &http.Client{Timeout: timeout}
	return &cli, nil
}
