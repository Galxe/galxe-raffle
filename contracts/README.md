# 🎲 Galxe Raffle Smart Contracts

## 📋 Requirements

- [Foundry](https://book.getfoundry.sh/getting-started/installation) for smart contract development

## 📄 Example Deployment

```sh
# Fmt Check
cd contracts/
forge fmt --check

# Build
cd contracts/
forge build

# Test
cd contracts/
forge test -vvv --match-contract '^RaffleTest$'
forge test -vvv --match-contract '^RaffleTest$' --match-test test_e2e_success

forge test -vvv --match-contract '^GGRaffleTest$'
forge test -vvv --match-contract '^GGRaffleTest$' --match-test test_e2e_success

# Deploy 
cd contracts/
forge script ./script/deploy/Raffle.s.sol:RaffleScript --broadcast --verify -vvvv

forge script ./script/deploy/GG_Raffle.s.sol:GGRaffleScript --broadcast --verify -vvvv

# Verify
cd contracts/
forge verify-contract 0x5F8A4108b1eD84512CD0FE766e18C93d105E89F0 \
    src/GG_Raffle.sol:GGRaffle \
    --verifier etherscan \
    --verifier-url https://explorer-sepolia.gravity.xyz/api \
    --etherscan-api-key "no-op" \
    --chain-id 13505
    
forge verify-contract 0x0e05faDf6713819FC8691fD686C58C241A0b517F \
    src/GG_Raffle.sol:GGRaffle \
    --verifier etherscan \
    --verifier-url https://explorer-sepolia.gravity.xyz/api \
    --etherscan-api-key "no-op" \
    --chain-id 13505
    
forge verify-contract 0xE95E06D5b7d80c8957754247eEab1E897f752783 \
    src/GG_Raffle.sol:GGRaffle \
    --verifier etherscan \
    --verifier-url https://rpc.gravity.xyz/api \
    --etherscan-api-key "no-op" \
    --chain-id 1625
    
forge verify-contract 0x0e05faDf6713819FC8691fD686C58C241A0b517F \
    src/GG_Raffle.sol:GGRaffle \
    --verifier etherscan \
    --verifier-url https://rpc.gravity.xyz/api \
    --etherscan-api-key "no-op" \
    --chain-id 1625
    
    
forge verify-contract --rpc-url https://rpc.gravity.xyz --verifier blockscout --verifier-url 'https://explorer-gravity-mainnet-0.t.conduit.xyz/api/' 0x7862c57E369c4035992aCdC5B5BcCEC0cfa3CDf4 src/GG_Raffle.sol:GGRaffle
    

forge verify-contract --rpc-url https://rpc.gravity.xyz --verifier blockscout --verifier-url 'https://explorer-gravity-mainnet-0.t.conduit.xyz/api/' 0x1f83DaDC3e56883E53Ed3370263A52dA7F4C19E1 src/Raffle.sol:Raffle

```

- Gravity Alpha Testnet Sepolia: [0x4349F1d73237c8Eb43FC04A16CFf70261D0F1343](https://explorer-sepolia.gravity.xyz/address/0x4349F1d73237c8Eb43FC04A16CFf70261D0F1343)
