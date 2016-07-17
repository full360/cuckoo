package consul

import consulapi "github.com/hashicorp/consul/api"

// Consul provides a client to the Consul API
type Consul struct {
	Client *consulapi.Client
}

// DefaultClientConfig returns the default consul config
func DefaultClientConfig() *consulapi.Config {
	return consulapi.DefaultConfig()
}

// NewClient creates and returns a consul client
func NewClient(config *consulapi.Config) (*Consul, error) {
	client, err := consulapi.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Consul{Client: client}, nil
}
