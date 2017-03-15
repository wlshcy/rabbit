.PHONY: build fmt 

# Prepend our vendor directory to the system GOPATH
# so that import path resolution will prioritize
# our third party snapshots.
export GO15VENDOREXPERIMENT=1
# GOPATH := ${PWD}/vendor:${GOPATH}
# export GOPATH

default: darwin 

darwin: fmt
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -v -o ./bin/rabbit ./src/

linux: fmt
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -o ./bin/rabbit ./src/

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./src/...

image:
	docker build -t wlsh/rabbit:latest .
