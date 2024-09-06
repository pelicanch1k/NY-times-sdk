package client

import (
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	return &Client{
		client: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (c *Client) GetAssets() (Answer, error) {

}
