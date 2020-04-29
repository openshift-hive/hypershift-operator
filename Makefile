IMG ?= hypershift-operator:latest

SRC_DIRS = cmd pkg

.PHONY: default
default: all

.PHONY: all
all: test build

.PHONY: build
build:  
	go build -mod=vendor -o bin/hypershift-operator github.com/openshift-hive/hypershift-operator/cmd/hypershift-operator

.PHONY: verify-gofmt
verify-gofmt:
	@echo Verifying gofmt
	@gofmt -l -s $(SRC_DIRS)>.out 2>&1 || true
	@[ ! -s .out ] || \
	  (echo && echo "*** Please run 'make fmt' in order to fix the following:" && \
	  cat .out && echo && rm .out && false)
	@rm .out

.PHONY: fmt
fmt:
	gofmt -w -s $(SRC_DIRS)

.PHONY: unit
unit:
	go test -mod=vendor ./pkg/... ./cmd/...

.PHONY: verify
verify: verify-gofmt

.PHONY: test
test: verify unit

.PHONY: build-container
build-container:
	buildah build-using-dockerfile --tag ${IMG} .

.PHONY: push-container
push-container: build-container
	buildah push ${IMG}
