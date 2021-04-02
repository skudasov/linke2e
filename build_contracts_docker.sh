#!/bin/bash
CONTRACTS_DIR=contracts_go_build
ETHEREUM_DOCKER_TOOLS_CONTAINER=ethereum/client-go:alltools-latest

# downloading modules
echo "downloading modules with yarn"
cd contracts_go_build && yarn install
cd ..

# abi's
echo "building abi's"
docker run --rm -v $(pwd)/$CONTRACTS_DIR:/root ethereum/solc:0.7.0 --abi /root/Greeter.sol --bin /root/Greeter.sol -o /root/abi --overwrite
docker run --rm -v $(pwd)/$CONTRACTS_DIR:/root ethereum/solc:0.6.6 --abi /root/APIConsumer.sol --bin /root/APIConsumer.sol -o /root/abi --overwrite
docker run --rm -v $(pwd)/$CONTRACTS_DIR:/root ethereum/solc:0.6.6 --abi /root/MockLink.sol --bin /root/MockLink.sol -o /root/abi --overwrite
docker run --rm -v $(pwd)/$CONTRACTS_DIR:/root ethereum/solc:0.6.6 --abi /root/MockOracle.sol --bin /root/MockOracle.sol -o /root/abi --overwrite

# stubs
echo "generating stubs"
mkdir -p $(pwd)/contracts_go_build/build
mkdir -p $(pwd)/contracts_go_build/build/greeter
mkdir -p $(pwd)/contracts_go_build/build/apiconsumer
mkdir -p $(pwd)/contracts_go_build/build/mocklink
mkdir -p $(pwd)/contracts_go_build/build/mockoracle

docker run --rm -v $(pwd)/$CONTRACTS_DIR:/root $ETHEREUM_DOCKER_TOOLS_CONTAINER abigen --abi /root/abi/Greeter.abi --bin /root/abi/Greeter.bin --pkg greeter --type Greeter --out /root/build/greeter/Greeter.go
docker run --rm -v $(pwd)/$CONTRACTS_DIR:/root $ETHEREUM_DOCKER_TOOLS_CONTAINER abigen --abi /root/abi/APIConsumer.abi --bin /root/abi/APIConsumer.bin --pkg apiconsumer --type APIConsumer --out /root/build/apiconsumer/APIConsumer.go
docker run --rm -v $(pwd)/$CONTRACTS_DIR:/root $ETHEREUM_DOCKER_TOOLS_CONTAINER abigen --abi /root/abi/MockLink.abi --bin /root/abi/MockLink.bin --pkg mocklink --type MockLink --out /root/build/mocklink/MockLink.go
docker run --rm -v $(pwd)/$CONTRACTS_DIR:/root $ETHEREUM_DOCKER_TOOLS_CONTAINER abigen --abi /root/abi/MockOracle.abi --bin /root/abi/MockOracle.bin --pkg mockoracle --type MockOracle --out /root/build/mockoracle/MockOracle.go



