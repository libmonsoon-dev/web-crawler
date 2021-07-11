package http

import (
	"net/http"
)

type ClientFactory interface {
	NewClient() *http.Client
}
