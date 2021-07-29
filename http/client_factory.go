package http

type ClientFactory interface {
	NewClient() *Client
}
