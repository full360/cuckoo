package main

import (
	"time"

	"github.com/full360/cuckoo/cloudwatch"
	"github.com/full360/cuckoo/consul"
	"github.com/full360/cuckoo/log"
)

// serviceCheckConfig is used to represent the configuration of a service check
type serviceCheckConfig struct {
	name            string
	tag             string
	metricName      string
	metricNamespace string
	blockTime       time.Duration
	logger          *log.Logger
}

// serviceCheck is used to represent a single service check with consul,
// cloudwatch and a logger
type serviceCheck struct {
	consul *consul.Check
	metric *cloudwatch.Metric
	logger *log.Logger
}

// defaultServiceCheck returns a defaul service check config
func defaultServiceCheck() *serviceCheckConfig {
	return &serviceCheckConfig{
		name:            "service",
		tag:             "tag",
		metricName:      "service_monitoring",
		metricNamespace: "microservices",
		blockTime:       10 * time.Minute,
		logger:          log.NewLogger(),
	}
}

// newServiceCheck returns a new service check
func newServiceCheck(svcConfig *serviceCheckConfig) (*serviceCheck, error) {
	consul, err := consul.NewCheck(&consul.CheckConfig{
		Service:     svcConfig.name,
		Tag:         svcConfig.tag,
		PassingOnly: true,
		BlockTime:   svcConfig.blockTime,
	})
	if err != nil {
		return nil, err
	}

	svcCheck := &serviceCheck{
		consul: consul,
		metric: cloudwatch.NewMetric(&cloudwatch.MetricConfig{
			Name:      svcConfig.metricName,
			Namespace: svcConfig.metricNamespace,
			Service: &cloudwatch.Service{
				Name: svcConfig.name,
				Env:  svcConfig.tag,
			},
			Value: 0,
		}),
		logger: svcConfig.logger,
	}
	return svcCheck, nil
}

// loopCheck does an infinite loop calling check
func (sc *serviceCheck) loopCheck() {
	for {
		err := sc.check()
		if err != nil {
			time.Sleep(10 * time.Second)
		}
	}
}

// check checks if a service is healthy and posts that data to a Cloudwatch
// metric based on the service name and environment
func (sc *serviceCheck) check() error {
	count, qm, err := sc.consul.Healthy()
	if err != nil {
		return err
	}
	// debug logging for Consul request
	sc.logger.Debug("Consul Query metadata, Request Time: %s, Last Index: %d", qm.RequestTime, qm.LastIndex)
	// Set the last response index as the wait index for the next request to
	// successfully do a blocking query
	sc.consul.QueryOptions.WaitIndex = qm.LastIndex
	sc.logger.Info("Service count: %d, with name: %s and tag: %s", count, sc.consul.Config.Service, sc.consul.Config.Tag)

	_, err = sc.metric.Put(float64(count))
	if err != nil {
		return err
	}
	return nil
}
