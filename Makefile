IMG ?= labring-action/sync:dev
TARGETARCH ?= amd64

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# only support linux, non cgo
PLATFORMS ?= linux_arm64 linux_amd64
GOOS=linux
CGO_ENABLED=0
GOARCH=$(shell go env GOARCH)

GO_BUILD_FLAGS=-trimpath -ldflags "-s -w"

.PHONY: all
all: build

.PHONY: clean
clean:
	rm -f bin/sync*

.PHONY: build
build: clean ## Build service-hub binary.
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) go build $(GO_BUILD_FLAGS) -o bin/sync main.go

.PHONY: docker-build
docker-build: build
	mv bin/sync bin/sync-${TARGETARCH}
	docker build -t $(IMG) . --build-arg TARGETARCH=${TARGETARCH}

.PHONY: docker-push
docker-push:
	docker push $(IMG)