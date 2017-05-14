#
# Makefile
# @author Evgeny Ukhanov <evgeny@ukhanov.org.ua>
#

.PHONY: test, build, fmt, install-go
default: build

GOVERSION = 1.8.1

test:
	@go test -v  $(go list ./... 2>&1 | grep -v "vendor")
	@go vet -v  $(go list ./... 2>&1 | grep -v "vendor")
	@echo $(gocover)

build:
	@go build

fmt:
	@gofmt -w -l .

install-go:
	@echo Installing Golang v$(GOVERSION)...
	@echo [================================================================]
	@wget -O /tmp/go.tar.gz https://storage.googleapis.com/golang/go$(GOVERSION).linux-amd64.tar.gz
	@sudo rm -f /usr/local/bin/go
	@sudo rm -f /usr/local/bin/godoc
	@sudo rm -f /usr/local/bin/gofmt
	@sudo tar -C /usr/local -xzf /tmp/go.tar.gz
	@sudo ln -s /usr/local/go/bin/go /usr/local/bin/go
	@sudo ln -s /usr/local/go/bin/godoc /usr/local/bin/godoc
	@sudo ln -s /usr/local/go/bin/gofmt /usr/local/bin/gofmt
	@echo [================================================================]
	@echo Done.
