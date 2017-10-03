package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/full360/cuckoo"
	"github.com/full360/cuckoo/log"
)

func main() {
	version := flag.Bool("version", false, "print version and exit")
	debug := flag.Bool("d", false, "enables debug logging mode")
	serviceName := flag.String("service", "", "Consul name of the Service to check")
	serviceTag := flag.String("tag", "", "Consul tag of the Service to check")
	metricName := flag.String("metric-name", "service_monitoring", "CloudWatch metric data name")
	metricNamespace := flag.String("metric-namespace", "microservices", "CloudWatch metric namespace")
	block := flag.Duration("block", 10*time.Minute, "Consul blocking query time")

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.Parse()
	if *version {
		fmt.Printf("cuckoo version: %s\n", Version)
		return
	}

	if *serviceName == "" {
		usageAndExit("Please enter a Service", 1)
	}

	if *serviceTag == "" {
		usageAndExit("Please enter a Service Tag", 1)
	}

	logger := log.NewLogger()

	if *debug {
		logger.SetLevel("debug")
	}

	svcCheck, err := cuckoo.NewServiceCheck(
		&cuckoo.ServiceCheckConfig{
			Name:            *serviceName,
			Tag:             *serviceTag,
			MetricName:      *metricName,
			MetricNamespace: *metricNamespace,
			BlockTime:       *block,
			Logger:          logger,
		},
	)
	if err != nil {
		logger.Fatal(fmt.Sprintf("%v", err))
	}

	for {
		err := svcCheck.Check()
		if err != nil {
			logger.Error(fmt.Sprintf("%v", err))
			time.Sleep(10 * time.Second)
		}
	}
}

// usageAndExit prints the default usage flags and exits the application with
// a status code (@jfrazelle code)
func usageAndExit(message string, exitCode int) {
	if message != "" {
		fmt.Fprintf(os.Stderr, message)
		fmt.Fprintf(os.Stderr, "\n\n")
	}
	flag.Usage()
	fmt.Fprintf(os.Stderr, "\n")
	os.Exit(exitCode)
}
