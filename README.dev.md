##### commands

Generate abi
```
docker run --rm -v $(pwd)/hardhat-template/contracts:/root -v $(pwd)/abi_build:/abi_build ethereum/solc:0.7.0 --abi /root/Greeter.sol -o ./abi_build --overwrite
```

Build golang client
```
docker run --rm -v $(pwd)/abi_build:/root ethereum/client-go:alltools-v1.10.1 abigen --abi /root/Greeter.abi --pkg greeter --type Greeter --out /root/Greeter.go
```

APIConsumer generate abi
```
docker run --rm -v $(pwd)/hardhat-template/contracts:/root -v $(pwd)/abi_build:/abi_build -v $(pwd)/hardhat-template/node_modules:/root/node_modules ethereum/solc:0.6.6 --abi /root/APIConsumer.sol -o ./abi_build --overwrite
```

Golang bindings generation from hardhat
```
cp -r hardhat-template/abi_export/ abi_build/
docker run --rm -v $(pwd)/hardhat-template/abi_export:/root ethereum/client-go:alltools-v1.10.1 abigen --abi /root/APIConsumer.json --pkg apiconsumer --type APIConsumer --out /root/APIConsumer.go
```

Manual building contracts with solc + abigen from docker
```
docker run --rm -v $(pwd)/contracts_go_build:/root ethereum/solc:0.7.0 --abi /root/Greeter.sol -o /root/abi --overwrite
docker run --rm -v $(pwd)/contracts_go_build:/root ethereum/client-go:alltools-v1.10.1 abigen --abi /root/abi/Greeter.abi --pkg greeter --type Greeter --out /root/build/Greeter.go

docker run --rm -v $(pwd)/contracts_go_build:/root ethereum/solc:0.6.6 --abi /root/APIConsumer.sol -o /root/abi --overwrite
docker run --rm -v $(pwd)/contracts_go_build:/root ethereum/client-go:alltools-v1.10.1 abigen --abi /root/abi/APIConsumer.abi --pkg apiconsumer --type APIConsumer --out /root/build/APIConsumer.go

docker run --rm -v $(pwd)/contracts_go_build:/root ethereum/solc:0.6.6 --abi /root/MockLink.sol -o /root/abi --overwrite
docker run --rm -v $(pwd)/contracts_go_build:/root ethereum/client-go:alltools-v1.10.1 abigen --abi /root/abi/MockLink.abi --pkg mocklink --type MockLink --out /root/build/MockLink.go

docker run --rm -v $(pwd)/contracts_go_build:/root ethereum/solc:0.6.6 --abi /root/MockOracle.sol -o /root/abi --overwrite
docker run --rm -v $(pwd)/contracts_go_build:/root ethereum/client-go:alltools-v1.10.1 abigen --abi /root/abi/MockOracle.abi --pkg mockoracle --type MockOracle --out /root/build/MockOracle.go
```

Clean up docker
```
docker system prune
docker volume rm $(docker volume ls -qf dangling=true)
docker rmi $(docker images | grep '^<none>' | awk '{print $3}')
```