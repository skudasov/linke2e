TEST_COUNT ?= 1
TEST_ARGS ?=
COVERPROFILE ?= coverage.out
BIN_DIR = bin
export GOPATH ?= $(shell go env GOPATH)
export GO111MODULE ?= on

.PHONY: lint
lint: golangci
	${BIN_DIR}/golangci-lint --color=always run ./... -v --timeout 5m

golangci: ## install golangci-linter
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${BIN_DIR} v1.27.0

go-acc:
	go get github.com/ory/go-acc@v0.2.3

install-deps: golangci go-acc

.PHONY: test
test:
	go test -v ./... -race -count $(TEST_COUNT) $(TEST_ARGS)

.PHONY: cover
cover: go-acc
	go-acc ./... && go tool cover -html=coverage.txt

.PHONY: start_geth
start_geth:
	cd deploy && docker-compose -f docker-compose-geth.yml up

.PHONY: stop_geth
stop_geth:
	cd deploy && docker-compose -f docker-compose-geth.yml down --remove-orphans

.PHONY: start_geth
start_hardhat:
	cd deploy && docker-compose -f docker-compose-hardhat.yml up

.PHONY: stop_geth
stop_hardhat:
	cd deploy && docker-compose -f docker-compose-hardhat.yml down --remove-orphans

.PHONY: build_contracts
build_contracts:
	./build_contracts_docker.sh
