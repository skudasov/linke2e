##### Chainlink node e2e tests

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