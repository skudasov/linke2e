## Chainlink node e2e tests

### Quickstart
Install deps, yarn is required
```
make install-deps
```
Start default geth development network
```
make start_geth
```
Run tests
```
make test_interactions
```

### Development (only OSX for now)

Setup deps
```
make install-deps
```

Start envs (hardhat is WIP for now)
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