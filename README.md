

<br/>`./build/bin/geth`
<br/>`/usr/local/bin/solc`
<br/>`/usr/local/bin/protoc`
<br/>`./cmd/abigen`

<br/>`solc --abi --bin ./contracts/GandalfV1ManagedFundFactory.sol -o build --overwrite`
<br/>`abigen --bin=./build/GandalfV1ManagedFundFactory.bin --abi=./build/GandalfV1ManagedFundFactory.abi --pkg=factory --out=./pkg/fund/managed/factory/GandalfV1ManagedFundFactory.go`
<br/>`abigen --bin=./build/GandalfV1ManagedSingleFund.bin --abi=./build/GandalfV1ManagedSingleFund.abi --pkg=single --out=./pkg/fund/managed/single/GandalfV1ManagedSingleFund.go`
<br/>`solc --abi --bin ./contracts/GandalfV1OracleList.sol -o build --overwrite`
<br/>`abigen --bin=./build/GandalfV1OracleList.bin --abi=./build/GandalfV1OracleList.abi --pkg=oraclelist --out=./pkg/oracle/list/GandalfV1OracleList.go`

```
go run main.go factorydeploy 1500000
go run main.go fundcreate "kovan test 0606-05" 1000000

go run main.go fundgettitle 0x7fB2E2D831e9B4eA48a2048ED6539Bf0581B3663
go run main.go fundgetlatestprice 0x7fB2E2D831e9B4eA48a2048ED6539Bf0581B3663 0x9b6Ff80Ff8348852d5281de45E66B7ED36E7B8a9
go run main.go fundgetoracleaddresses 0x7fB2E2D831e9B4eA48a2048ED6539Bf0581B3663

go run main.go fundswap 0x2058f8C98e3f2d88B58FB8162f84c58dD489441b

go run main.go fundtest 0xC8c8fa2695610aeBb309580059f764B4730FF7a2
```

```
truffle network
```