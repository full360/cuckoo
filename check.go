package main

import (
	"time"

	"github.com/full360/health/cloudwatch"
	"github.com/full360/health/consul"
	"github.com/full360/health/log"
)

// Check is used to represent a single service check with consul, cloudwatch
// and a logger
type Check struct {
	consul *consul.Check
	metric *cloudwatch.Metric
	logger *log.Logger
}

// NewServiceCheck returns a new service check
func NewServiceCheck(name, tag string, block time.Duration, logger *log.Logger) (*Check, error) {
	consulConfig := consul.DefaultCheckConfig()
	config := cloudwatch.DefaultMetricConfig()
	if name != "" {
		consulConfig.Service = name
		config.Service.Name = name
	}
	if tag != "" {
		consulConfig.Tag = tag
		config.Service.Env = tag
	}
	consulConfig.BlockTime = block

	metric := cloudwatch.NewMetric(config)

	consul, err := consul.NewCheck(consulConfig)
	if err != nil {
		return nil, err
	}

	serviceCheck := &Check{
		consul: consul,
		metric: metric,
		logger: logger,
	}
	return serviceCheck, nil
}

// LoopServiceCheck does an infinite loop calling serviceCheck
func (c *Check) LoopServiceCheck() {
	for {
		err := c.serviceCheck()
		if err != nil {
			time.Sleep(10 * time.Second)
		}
	}
}

// serviceCheck checks if a service is healthy and posts that data to a
// Cloudwatch metric based on the service name and environment
func (c *Check) serviceCheck() error {
	count, qm, err := c.consul.Healthy()
	if err != nil {
		c.logger.Error("Could not retrieve service count from Consul: %v", err)
		return err
	}
	// debug logging for Consul request
	c.logger.Debug("Consul Query metadata, Request Time: %s, Last Index: %d", qm.RequestTime, qm.LastIndex)
	// Set the last response index as the wait index for the next request to
	// successfully do a blocking query
	c.consul.QueryOptions.WaitIndex = qm.LastIndex
	c.logger.Info("Service count: %d, with name: %s and tag: %s", count, c.consul.Service, c.consul.Tag)

	_, err = c.metric.Put(float64(count))
	if err != nil {
		c.logger.Error("Could not post metric to CloudWatch: %v", err)
		return err
	}
	return nil
}
