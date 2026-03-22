DIST = dist
GOCMD = go
GOBUILD = $(GOCMD) build
GOMOD = $(GOCMD) mod
GOTEST = $(GOCMD) test
GITHASH = $(shell git rev-parse HEAD)
BUILDDATE = $(shell date "+%F_%H:%M:%S%z")
VERSION = $(shell cat VERSION)
LD_FLAGS = -X 'pingcode-client/info.Version=${VERSION}' -X 'pingcode-client/info.GitHash=${GITHASH}' -X 'pingcode-client/info.BuildDate=${BUILDDATE}'
BIN_NAME = pingcode-client

.PHONY: install clean run test build-linux-amd64 build-darwin-amd64 build-linux-arm64 build-darwin-arm64 build-linux build-darwin build

install:
	$(GOMOD) tidy

clean:
	@$(GOCMD) clean
	rm -rf ${DIST}/*

run: install
	@$(GOCMD) run -ldflags="${LD_FLAGS}" main.go

test: install
	@$(GOTEST) -ldflags="${LD_FLAGS}" -v ./...

build-linux-amd64:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux $(GOBUILD) -ldflags="${LD_FLAGS}" -o ${DIST}/${BIN_NAME}.amd64 main.go 

build-darwin-amd64:
	CGO_ENABLED=1 GOARCH=amd64 GOOS=darwin $(GOBUILD) -ldflags="${LD_FLAGS}" -o ${DIST}/${BIN_NAME}.darwin-amd64 main.go

build-linux-arm64:
	CGO_ENABLED=0 GOARCH=arm64 GOOS=linux $(GOBUILD) -ldflags="${LD_FLAGS}" -o ${DIST}/${BIN_NAME}.arm64 main.go 

build-darwin-arm64:
	CGO_ENABLED=1 GOARCH=arm64 GOOS=darwin $(GOBUILD) -ldflags="${LD_FLAGS}" -o ${DIST}/${BIN_NAME}.darwin-arm64 main.go

build-linux: build-linux-amd64 build-linux-arm64

build-darwin: build-darwin-amd64 build-darwin-arm64

build: install build-linux build-darwin
