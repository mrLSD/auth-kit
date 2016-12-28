#
# Makefile
# @author Evgeny Ukhanov <mrlsd@ya.ru>
#

.PHONY: run, test, build, fmt
default: run

GOVERSION = 1.7.3

run:
	@go run server.go

test:
	@go test -v  $(go list ./... 2>&1 | grep -v "vendor")
	@go vet -v  $(go list ./... 2>&1 | grep -v "vendor")

build:
	@go build

fmt:
	@gofmt -w -l .

install-go:
	@echo Installing Golang v$(GOVERSION)...
	@echo [================================================================]
	@wget -O /tmp/go.tar.gz https://storage.googleapis.com/golang/go$(GOVERSION).linux-amd64.tar.gz
	@sudo tar -C /usr/local -xzf /tmp/go.tar.gz
	@sudo ln -s /usr/local/go/bin/go /usr/local/bin/go
	@sudo ln -s /usr/local/go/bin/godoc /usr/local/bin/godoc
	@sudo ln -s /usr/local/go/bin/gofmt /usr/local/bin/gofmt
	@echo [================================================================]
	@echo Done.

rerun-install:
	@go get github.com/tockins/realize

rerun:
	@realize run