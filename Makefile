#-include ./vendor/lenovo.com/aimaster-k8s-common/pkg/version/ldflags
#OUTPUT=./bin/letrain-cv-train-service
OS=linux
ARCH=amd64
SRC=./cmd

all: format build

vendor:
	go mod vendor

build:
	GOOS=$(OS) GOARCH=$(ARCH) go build -ldflags "-s -w ${LDFLAGS}" -o $(OUTPUT) $(SRC)

vet:
	go vet ./...

test: vet
	go test -v ./...

clean:
	rm -rf ./$(OUTPUT)

.PHONY: format build clean vet test
