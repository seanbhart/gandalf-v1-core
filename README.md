## NOTE: This repo is incomplete.

- Solidity contracts contain hard-coded Ethereum addresses and unfinished / missing functions.
- Go handler functions contain incomplete error handling and missing contract control functions.
- Test modules are missing.

---

# Gandalf DeFi Asset Management

Gandalf V1 was originally working toward the creation of modularized asset management tools. This first draft includes three incomplete components of this protocol:

- A fund contract that holds deposited assets and allows a single manager to swap assets with other whitelisted assets
- A factory contract that creates new single manager fund contracts upon request
- An oracle contract that looks up external prices to ensure the proper amount of tokens are exchanged during swaps

## Go Setup

ABIs (Application Binary Interfaces) are required to interact with an Ethereum contract in Go. See the Go Ethereum Book's [Smart Contract Compilation & ABI](https://goethereumbook.org/en/smart-contract-compile/) page for installation instructions for the Solidity compiler and ABI Generator tools. The `abigen` tool will also convert the ABI file to a Go file for import, allowing referencing of the contract functions directly. This is a much cleaner way to reference the functions rather than using `delegatecall`.

## Solidity Compilation | ABI & Go File Generation

Compile both the Factory contract and the Oracle contract. The Fund contract will be compiled when the Factory contract is compiled because it is referenced by the Fund Factory.

```
solc --abi --bin ./contracts/GandalfV1ManagedFundFactory.sol -o build --overwrite
solc --abi --bin ./contracts/GandalfV1OracleList.sol -o build --overwrite
```

Generate all ABI and Go files from the three contracts.

```
abigen --bin=./build/GandalfV1ManagedFundFactory.bin --abi=./build/GandalfV1ManagedFundFactory.abi --pkg=factory --out=./pkg/fund/managed/factory/GandalfV1ManagedFundFactory.go
abigen --bin=./build/GandalfV1ManagedSingleFund.bin --abi=./build/GandalfV1ManagedSingleFund.abi --pkg=single --out=./pkg/fund/managed/single/GandalfV1ManagedSingleFund.go
abigen --bin=./build/GandalfV1OracleList.bin --abi=./build/GandalfV1OracleList.abi --pkg=oraclelist --out=./pkg/oracle/list/GandalfV1OracleList.go
```

## Network Setup

Create an `.env` file and populate it with data in the following format.

Data notes:

- `ACCOUNT_KEY_PRIV` should be the account that also owns the contracts so all tests can be used.
- `NETWORK` can be a local network, a local node port, or a third party network (e.g. infura).
- `FACTORY_ADDRESS` is the address of the Factory contract (once deployed) on the listed network
- `ORACLE_ADDRESS` is the address of the Oracle contract (once deployed) on the listed network

```
ACCOUNT_KEY_PRIV=
NETWORK=
FACTORY_ADDRESS=
ORACLE_ADDRESS=
```

## Running Go contract handler functions

Note: Customize the command portions in brackets. These example tests are not comprehensive.

1. Deploy the Factory Contract

   - e.g.: `go run main.go factorydeploy [2000000]`
   - The transaction hash will be the last output line. Look up the transaction on etherscan if using a testnet. (e.g. `https://kovan.etherscan.io/tx/0x0c933794f929f96624011f9832f93b8cd57e24d53fbce886e602ef91bdf44525`)
   - The Factory contract address will be the second output line. Save this address for the third step.

2. Deploy the Oracle Contract

   - e.g.: `go run main.go oracledeploy [750000]`
   - The Oracle contract address will be the second output line. Save this address for the next step.

3. Save the Factory and Oracle contract addresses in the `.env` file.

   - Edit the `FACTORY_ADDRESS=` to store the new Factory address
   - Edit the `ORACLE_ADDRESS=` to store the new Oracle address

4. Create a new Fund contract via the Factory contract.

   - e.g.: `go run main.go fundcreate ["kovan test 1003-01"] [0x18b20Af6ad00eb73de9Af3d12e2698dAd23ad4bf] [1500000]`
   - The transaction hash will be the last output line. Look up the transaction and ensure it did not run out of gas.

5. Get the new Fund address by querying the latest Fund by index.

   - e.g.: `go run main.go fundqueryindex [0]`
   - The Fund address will be the last output line. Save the address for the next step.

6. Check the Fund title.
   - e.g.: `go run main.go fundgettitle [0x7498C732f922b662DA564b6D096A370332203918]`
   - e.g. for the price of WETH on Kovan: `go run main.go fundgetlatestprice [0x7498C732f922b662DA564b6D096A370332203918] [0xd0A1E359811322d97991E03f863a0C30C2cF029C]`

## Final Note

Thanks for stopping by! Gandalf has progressed to other DeFi solutions, so this repo might not ever be completed, but hopefully it provided a small example of the potential for Go-based Ethereum contract handling.
