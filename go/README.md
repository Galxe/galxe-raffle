# 🎲 Galxe Raffle

This directory contains the Go bindings for the Galxe Raffle smart contract, as well as a library for running the raffle offchain.

## 📄 Binding

To generate the binding, run the following command:

```sh
cat ../contracts/out/IRaffle.sol/IRaffle.json | jq -r .abi > abi
abigen --abi abi --type raffle --out pkg/contract/binding.go --pkg contract
```
