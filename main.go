package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/full360/health/log"
)

var (
	version bool
	debug   bool
	service string
	block   time.Duration
	tag     string
)

func init() {
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&debug, "d", false, "enables debug logging mode")
	flag.StringVar(&service, "service", "", "Name of the Service to check")
	flag.StringVar(&tag, "tag", "", "Tag name of the Service to check")
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

	if service == "" {
		usageAndExit("Please enter a Service", 0)
	}

	if tag == "" {
		usageAndExit("Please enter a Service Tag", 0)
	}

	logger := log.NewLogger()

	if debug {
		logger.SetLevel("debug")
	}

	check, err := NewServiceCheck(service, tag, block, logger)
	if err != nil {
		logger.Fatal("%v", err)
	}
	check.LoopServiceCheck()
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
