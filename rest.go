package restapi

import (
	"github.com/go-zoox/restapi/client"
	"github.com/go-zoox/restapi/server"
)

// NewClient creates a REST API client.
func NewClient(uri string, cfg ...*client.Config) *client.Client {
	return client.New(uri, cfg...)
}

// NewServer creates a REST API server.
func NewServer(port int, cfg ...*server.Config) *server.Server {
	return server.New(port, cfg...)
}
