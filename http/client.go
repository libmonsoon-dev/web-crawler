package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/libmonsoon-dev/web-crawler/http/useragent"
)

func NewClient(client *http.Client) *Client {
	return &Client{Client: client}
}

type Client struct {
	*http.Client

	mu        sync.Mutex
	userAgent string
}

func (c *Client) Get(ctx context.Context, url string) (resp *http.Response, err error) {
	req, err := c.newRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) GetUserAgent() string {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.userAgent == "" {
		c.userAgent = useragent.Computer()
	}

	return c.userAgent
}

func (c *Client) SetUserAgent(ua string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.userAgent = ua
}

func (c *Client) newRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("User-Agent", c.GetUserAgent())
	req = req.WithContext(ctx)
	return req, nil
}
