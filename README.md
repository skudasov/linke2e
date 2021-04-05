## Chainlink node e2e tests

### Development (only OSX for now)

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

## Miscellaneous
### Outgoing call to mocks
Mock API runs on host network, you need to call host.docker.internal from docker, example:
```
http://host.docker.internal:9092
```
Check your DNS from inside
```
docker run --rm alpine nslookup host.docker.internal
```
Use default Docker Engine config, without additional dns
```
{
  "debug": true,
  "experimental": false
}
```