package pld

import "fmt"

type Client struct {
	name string
}

func (c *Client) SetName(name string) {
	c.name = name
}

func (c *Client) GetName() string {
	return c.name
}

func (c *Client) PrintName() {
	fmt.Println(c.name)
}
