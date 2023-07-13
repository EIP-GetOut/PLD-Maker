package gdoc

import (
	"encoding/json"
)

type Client struct {
	Left int
}

func NewClient(conf []byte) (*Client, error) {
	var cli Client

	if err := json.Unmarshal(conf, &cli); err != nil {
		return nil, err
	}
	return &cli, nil
}
