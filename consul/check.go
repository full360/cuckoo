package consul

import (
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

// CheckConfig is used to represent a basic Check configuration
type CheckConfig struct {
	// Service the name of the service we'll check
	Service string
	// Tags is a filter applied to the service we'll check
	Tag string
	// PassingOnly will show only passing services or not
	PassingOnly bool
	// BlockTime Consul blocking query time
	BlockTime time.Duration
}

// Check is used to represent a single service check
type Check struct {
	// Consul Client object from our custom Consul client wrapper
	Consul *Consul
	// CheckConfig
	Config *CheckConfig
	// QueryOptions is a consulapi struct that allows us to to blocking queries
	QueryOptions *consulapi.QueryOptions
}

// DefaultCheckConfig retruns a valid check configuration
func DefaultCheckConfig() *CheckConfig {
	return &CheckConfig{
		Service:     "default_service",
		Tag:         "default_tag",
		PassingOnly: true,
		BlockTime:   10 * time.Minute,
	}
}

// NewCheck returns a new check
func NewCheck(config *CheckConfig) (*Check, error) {
	consulClient, err := NewClient(DefaultClientConfig())
	if err != nil {
		return nil, err
	}

	check := &Check{
		Consul: consulClient,
		Config: config,
		QueryOptions: &consulapi.QueryOptions{
			WaitIndex: 0,
			WaitTime:  config.BlockTime,
		},
	}
	return check, nil
}

// Healthy returns the number of healthy instances of a service based on the
// data of the exisitng check
func (c *Check) Healthy() (int, *consulapi.QueryMeta, error) {
	out, qm, err := c.Consul.Client.Health().Service(c.Config.Service, c.Config.Tag, c.Config.PassingOnly, c.QueryOptions)
	if err != nil {
		return -1, nil, err
	}
	return len(out), qm, nil
}
