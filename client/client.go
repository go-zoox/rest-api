package client

// Client is a JSONRPC client.
type Client struct {
	Endpoint string
	//
	Config *Config
}

// Config is a JSONRPC client configuration.
type Config struct {
	Headers map[string]string
	Query   map[string]string
}

// New creates a JSONRPC client.
func New(uri string, cfg ...*Config) *Client {
	var config *Config
	if len(cfg) > 0 && cfg[0] != nil {
		config = cfg[0]
	} else {
		config = &Config{}
	}

	return &Client{
		Endpoint: uri,
		Config:   config,
	}
}

// // @TODO method must have no type parameterssyntax
// // CreateResource creates a rest api resource
// func (c *Client) CreateResource[T any](namespace string) *Resource[T] {
// 	return &Resource[T]{
// 		client:    c,
// 		namespace: namespace,
// 	}
// }

// CreateResource creates a rest api resource
func CreateResource[T any](c *Client, namespace string) *Resource[T] {
	return &Resource[T]{
		client:    c,
		namespace: namespace,
	}
}
