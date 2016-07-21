GOTOOLS=github.com/mitchellh/gox
PACKAGES=$(shell go list ./... | grep -v '^github.com/full360/cuckoo/vendor/')
VETARGS?=-asmdecl -atomic -bool -buildtags -copylocks -methods \
         -nilfunc -printf -rangeloops -shift -structtags -unsafeptr
VERSION?=$(shell awk -F\" '/^const Version/ { print $$2; exit }' version.go)

# builds and generate package distributions
all: build

# dev creates binaries for testing locally - these are put into ./bin and $GOPATH
# NOTE: As we need to connect to a VPC using DNS and the golang dns does not
# picks up viscosity dns injections it will not work without CGO_ENABLED=1
local: format
	@LOCAL_BUILD=1 sh -c "'$(CURDIR)/scripts/build.sh'"

# all builds binaries for all targets
build: tools
	@mkdir -p bin/
	@sh -c "'$(CURDIR)/scripts/build.sh'"

# dist builds binaries for all platforms and packages them for distribution
dist: build
	@sh -c "'$(CURDIR)/scripts/dist.sh' $(VERSION)"

format:
	@echo "--> Running go fmt"
	@go fmt $(PACKAGES)

tools:
	go get -u -v $(GOTOOLS)

vet:
	@echo "--> Running go tool vet $(VETARGS) ."
	@go list ./... \
		| grep -v ^github.com/full360/cuckoo/vendor/ \
		| cut -d '/' -f 4- \
		| xargs -n1 \
			go tool vet $(VETARGS) ;

test: format
	@$(MAKE) vet
	@sh -c "'$(CURDIR)/scripts/test.sh'"

docker:
	docker build -t full360/cuckoo:latest . \
	&& docker tag -f full360/cuckoo:latest full360/cuckoo:$(VERSION) \
	&& docker push full360/cuckoo

.PHONY: all local build dist format tools vet test docker
