##### Chainlink node e2e tests

#### Development

Setup deps
```
go mod download
cd hardhat-template
yarn install
```

Start envs
```
make start_geth
make stop_geth

make start_hardhat
make stop_hardhat
```
Run golang lint
```
make lint
```
Rebuild contracts (golang deployment)
```
./build_contracts_docker.sh
```

Run tests
```
make test_interactions
```