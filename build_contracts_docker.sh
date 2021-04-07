#!/bin/bash
CONTRACTS_DIR=contracts_go_build
ETHEREUM_DOCKER_TOOLS_CONTAINER=ethereum/client-go:alltools-latest
DEFAULT_SOLC_CONTAINER=ethereum/solc:0.6.0

echo "downloading modules with yarn"
cd contracts_go_build && yarn install
cd ..

# shellcheck disable=SC2207
# shellcheck disable=SC2010
FILES=($(ls $CONTRACTS_DIR | grep '.sol' | sed -e 's/\..*$//'))
mkdir -p "$(pwd)"/contracts_go_build/build

for i in "${FILES[@]}"
do
  echo "building abi's for ${i}"
  docker run --rm -v "$(pwd)"/$CONTRACTS_DIR:/root "${DEFAULT_SOLC_CONTAINER}" --abi /root/"${i}".sol --bin /root/"${i}".sol -o /root/abi --overwrite
  mkdir -p "$(pwd)"/contracts_go_build/build/"${i}"
  echo "generating stubs for ${i}"
  docker run --rm -v "$(pwd)"/$CONTRACTS_DIR:/root $ETHEREUM_DOCKER_TOOLS_CONTAINER abigen --abi /root/abi/"${i}".abi --bin /root/abi/"${i}".bin --pkg "${i}" --type "${i}" --out /root/build/"${i}"/"${i}".go
done



