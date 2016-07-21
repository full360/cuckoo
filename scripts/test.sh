#!/usr/bin/env bash

# This script was taken from Hashicorp Consul and modified to feed this project
# needs. https://github.com/hashicorp/consul/blob/master/scripts/test.sh
# Create a temp dir and clean it up on exit
TEMPDIR=`mktemp -d -t cuckoo-test.XXX`
trap "rm -rf $TEMPDIR" EXIT HUP INT QUIT TERM

# Build the Cuckoo binary for the API tests
echo "--> Building Cuckoo"
go build -o $TEMPDIR/cuckoo || exit 1

# Run the tests
echo "--> Running tests"
go list ./... | grep -v '^github.com/full360/cuckoo/vendor/' | PATH=$TEMPDIR:$PATH xargs -n1 go test ${GOTEST_FLAGS:--cover -timeout=360s}
