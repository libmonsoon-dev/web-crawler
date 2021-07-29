package http

import (
	"context"
	"net/http"
)

func NewClient(client *http.Client) *Client {
	return &Client{client}
}

type Client struct {
	*http.Client
}

func (c *Client) Get(ctx context.Context, url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	return c.Do(req)
}
