SHELL=bash

GOOS?=linux
GOARCH?=amd64

build: publish-log publish-performance

publish-log:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o log-metrics publish-log-analysis/main.go

publish-performance:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o performance publishing-performance-analysis/main.go


.PHONEY: test build debug
