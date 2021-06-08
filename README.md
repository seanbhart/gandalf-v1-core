

<br/>`./build/bin/geth`
<br/>`/usr/local/bin/solc`
<br/>`/usr/local/bin/protoc`
<br/>`./cmd/abigen`

<br/>`solc --abi --bin ./contracts/GandalfV1ManagedFundFactory.sol -o build --overwrite`
<br/>`abigen --bin=./build/GandalfV1ManagedFundFactory.bin --abi=./build/GandalfV1ManagedFundFactory.abi --pkg=factory --out=./pkg/fund/managed/factory/GandalfV1ManagedFundFactory.go`
<br/>`abigen --bin=./build/GandalfV1ManagedSingleFund.bin --abi=./build/GandalfV1ManagedSingleFund.abi --pkg=single --out=./pkg/fund/managed/single/GandalfV1ManagedSingleFund.go`
<br/>`solc --abi --bin ./contracts/GandalfV1OracleList.sol -o build --overwrite`
<br/>`abigen --bin=./build/GandalfV1OracleList.bin --abi=./build/GandalfV1OracleList.abi --pkg=oraclelist --out=./pkg/oracle/list/GandalfV1OracleList.go`