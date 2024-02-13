package observer

import (
	"context"
	"fmt"
)

type Client struct {
	name string
	age  int
}

func (c *Client) Update(ctx context.Context) {
	fmt.Printf("Client %s (%d y.o.)got update about end of training\n", c.name, c.age)
}

func (c *Client) UpdateName(name string) {
	c.name = name
}

func NewClient(name string, age int) *Client {
	return &Client{
		name: name,
		age:  age,
	}
}
