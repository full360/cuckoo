package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/full360/health/log"
)

var (
	version         bool
	debug           bool
	serviceName     string
	serviceTag      string
	metricName      string
	metricNamespace string
	block           time.Duration
)

func init() {
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&debug, "d", false, "enables debug logging mode")
	flag.StringVar(&serviceName, "service", "", "Consul name of the Service to check")
	flag.StringVar(&serviceTag, "tag", "", "Consul tag of the Service to check")
	flag.StringVar(&metricName, "metric-name", "service_monitoring", "CloudWatch metric data name")
	flag.StringVar(&metricNamespace, "metric-namespace", "microservices", "CloudWatch metric namespace")
	flag.DurationVar(&block, "block", 10*time.Minute, "Consul blocking query time")

	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.Parse()
}

func main() {
	if version {
		fmt.Printf("health version: %s\n", Version)
		return
	}

	if serviceName == "" {
		usageAndExit("Please enter a Service", 0)
	}

	if serviceTag == "" {
		usageAndExit("Please enter a Service Tag", 0)
	}

	logger := log.NewLogger()

	if debug {
		logger.SetLevel("debug")
	}

	svcCheck, err := newServiceCheck(&serviceCheckConfig{
		name:            serviceName,
		tag:             serviceTag,
		metricName:      metricName,
		metricNamespace: metricNamespace,
		blockTime:       block,
		logger:          logger,
	})
	if err != nil {
		logger.Fatal(fmt.Sprintf("%v", err))
	}
	svcCheck.loopCheck()
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
