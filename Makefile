#!/usr/bin/make -f

PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed 's/ /\@/g')
BINDIR ?= $(GOPATH)/bin
SIMAPP = ./app

# for dockerized protobuf tools
DOCKER := $(shell which docker)
BUF_IMAGE=bufbuild/buf #@sha256:3cb1f8a4b48bd5ad8f09168f10f607ddc318af202f5c057d52a45216793d85e5 #v1.4.0
DOCKER_BUF := $(DOCKER) run --platform="linux/amd64" --rm -v $(CURDIR):/workspace --workdir /workspace $(BUF_IMAGE)
HTTPS_GIT := https://github.com/jackalLabs/canine-chain.git

export GO111MODULE = on

###############################################################################
###                          Go Toolchain Handling                          ###
###############################################################################
# This block extracts the desired toolchain version from go.mod (e.g., "go1.23.1").
# It then displays both the current Go version and the expected one. If the expected
# version is not already active (or in PATH), it installs it using golang.org/dl.
#
# Note: The toolchain version is expected to be defined exactly as the desired binary
# name (e.g., "go1.23.1"). This section does not alter the go.mod syntax.
CURRENT_GO_VERSION := $(shell go version)
$(info Current Golang Version: $(CURRENT_GO_VERSION))

TOOLCHAIN_GO_VERSION := $(shell grep '^toolchain' go.mod | awk '{print $$2}')
ifneq ($(TOOLCHAIN_GO_VERSION),)
  $(info Wanted Golang Version as specified in go.mod: $(TOOLCHAIN_GO_VERSION))
  ifeq ("$(shell which $(TOOLCHAIN_GO_VERSION) 2>/dev/null)","")
    ifeq ("$(shell go version | grep $(TOOLCHAIN_GO_VERSION))","")
      $(info The wanted Golang version was not found. Installing $(TOOLCHAIN_GO_VERSION) via golang.org/dl...)
      $(shell go install golang.org/dl/$(TOOLCHAIN_GO_VERSION)@latest)
      $(shell $(TOOLCHAIN_GO_VERSION) download)
    else
      $(info The wanted Golang version is already active.)
    endif
  else
    $(info Found $(TOOLCHAIN_GO_VERSION) in PATH.)
  endif
  GO_CMD := $(TOOLCHAIN_GO_VERSION)
else
  $(info No toolchain version specified in go.mod. Using default 'go' command.)
  GO_CMD := go
endif

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif

###############################################################################
###                          PebbleDB Opt-In Logic                          ###
###############################################################################
# If WITH_PEBBLEDB=yes is provided, we enable the 'pebbledb' build tag, add
# necessary ldflags for Pebble, and append "-pebbledb" to VERSION.
# (Note: if you require a separate go.mod for PebbleDB, handle that externally.)
ifeq ($(WITH_PEBBLEDB),yes)
  build_tags += pebbledb
  VERSION := $(VERSION)-pebbledb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=pebbledb -X github.com/tendermint/tm-db.ForceSync=1
  $(info Applying PebbleDB support. Make sure to use an alternate go.mod (e.g., go-4pebbledb.mod) if needed.)
endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
empty = $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(empty),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=canine \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=canined \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X github.com/jackalLabs/canine-chain/app.Bech32Prefix=jkl \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif

ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags_comma_sep)" -ldflags '$(ldflags)' -trimpath

# The below include contains the tools and runsim targets.
# include contrib/devtools/Makefile

all: install lint test

build: go.sum
ifeq ($(OS),Windows_NT)
	exit 1
else
	$(GO_CMD) build -mod=readonly $(BUILD_FLAGS) -o build/canined ./cmd/canined
endif

build_cli: build

build-contract-tests-hooks:
ifeq ($(OS),Windows_NT)
	$(GO_CMD) build -mod=readonly $(BUILD_FLAGS) -o build/contract_tests.exe ./cmd/contract_tests
else
	$(GO_CMD) build -mod=readonly $(BUILD_FLAGS) -o build/contract_tests ./cmd/contract_tests
endif

install: go.sum
	$(GO_CMD) install -mod=readonly $(BUILD_FLAGS) ./cmd/canined

########################################
### Tools & dependencies
########################################

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@$(GO_CMD) mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@$(GO_CMD) mod verify

draw-deps:
	@# requires brew install graphviz or apt-get install graphviz
	$(GO_CMD) get github.com/RobotsAndPencils/goviz
	@goviz -i ./cmd/canined -d 2 | dot -Tpng -o dependency-graph.png

clean:
	rm -rf snapcraft-local.yaml build/

distclean: clean
	rm -rf vendor/

########################################
### Testing
########################################

local: install
	./scripts/test-node.sh $(address)

test: test-unit
test-all: test-race test-cover
test-sim: test-sim-import-export test-sim-full-app

test-unit:
	@VERSION=$(VERSION) $(GO_CMD) test -short -mod=readonly -tags='ledger test_ledger_mock' ./...

test-race:
	@VERSION=$(VERSION) $(GO_CMD) test -mod=readonly -race -tags='ledger test_ledger_mock' ./...

test-cover:
	@$(GO_CMD) test -mod=readonly -timeout 30m -race -coverprofile=coverage.txt -covermode=atomic -tags='ledger test_ledger_mock' ./...

benchmark:
	@$(GO_CMD) test -mod=readonly -bench=. ./...

test-sim-import-export: runsim
	@echo "Running application import/export simulation. This may take several minutes..."
	@runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 50 5 TestAppImportExport

test-sim-full-app: runsim
	@echo "Running short multi-seed application simulation. This may take awhile!"
	@runsim -Jobs=4 -SimAppPkg=$(SIMAPP) -ExitOnFail 50 10 TestFullAppSimulation

test-sim-bench:
	@VERSION=$(VERSION) $(GO_CMD) test -benchmem -run ^BenchmarkFullAppSimulation -bench ^BenchmarkFullAppSimulation -cpuprofile cpu.out github.com/jackalLabs/canine-chain/app

runsim:
	$(GO_CMD) install github.com/cosmos/tools/cmd/runsim@latest

###############################################################################
###                                Linting                                  ###
###############################################################################

format-tools:
	$(GO_CMD) install mvdan.cc/gofumpt@v0.6.0
	gofumpt -l -w .

lint: format-tools
	golangci-lint run --fix

format: format-tools
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs gofumpt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/lcd/statik/statik.go" | xargs goimports -w -local github.com/jackalLabs/canine-chain

###############################################################################
###                                Protobuf                                 ###
###############################################################################
# thanks juno ;)
protoVer=v0.7
protoImageName=tendermintdev/sdk-proto-gen:$(protoVer)
containerProtoGen=jackal-proto-gen-$(protoVer)
containerProtoGenAny=jackal-proto-gen-any-$(protoVer)
containerProtoGenSwagger=jackal-proto-gen-swagger-$(protoVer)
containerProtoFmt=jackal-proto-fmt-$(protoVer)

proto-all: proto-format proto-lint proto-gen

proto-gen:
	@echo "Generating Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/protocgen.sh; fi

# This generates the SDK's custom wrapper for google.protobuf.Any. It should only be run manually when needed
proto-gen-any:
	@echo "Generating Protobuf Any"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGenAny}$$"; then docker start -a $(containerProtoGenAny); else docker run --name $(containerProtoGenAny) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/protocgen-any.sh; fi

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGenSwagger}$$"; then docker start -a $(containerProtoGenSwagger); else docker run --name $(containerProtoGenSwagger) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/protoc-swagger-gen.sh; fi

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \; ; fi

proto-lint:
	@$(DOCKER_BUF) lint --error-format=json

proto-check-breaking:
	@$(DOCKER_BUF) breaking --against $(HTTPS_GIT)#branch=main

# Note: The following targets are declared in .PHONY but have no corresponding rules:
#       - install-debug, test-build, proto-update-deps
# They are placeholders for future enhancements.

.PHONY: proto-all proto-gen proto-gen-any proto-swagger-gen proto-format proto-lint proto-check-breaking proto-update-deps docs
.PHONY: all install install-debug go-mod-cache draw-deps clean build format test test-all test-build test-cover test-unit test-race test-sim-import-export local
