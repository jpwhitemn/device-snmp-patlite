.PHONY: build test clean prepare update

GO=CGO_ENABLED=0 go

MICROSERVICES=cmd/device-snmp-patlite
.PHONY: $(MICROSERVICES)

VERSION=$(shell cat ./VERSION)

GOFLAGS=-ldflags "-X github.com/edgexfoundry/device-snmp-patlite.Version=$(VERSION)"

build: $(MICROSERVICES)
	go build ./...

cmd/device-snmp-patlite:
	$(GO) build $(GOFLAGS) -o $@ ./cmd

test:
	go test ./... -cover

clean:
	rm -f $(MICROSERVICES)

prepare:
	glide install

update:
	glide update
