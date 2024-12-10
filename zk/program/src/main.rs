//! A program that selects winners for a raffle and generates a Merkle root for winners

#![no_main]
sp1_zkvm::entrypoint!(main);

use alloy_sol_types::sol;
use alloy_sol_types::SolType;
use raffle_lib::raffle;
use sha2_v0_9_8::{Digest as Digest_9_8, Sha256 as Sha256_9_8};

pub fn main() {
    // Read the raffle event data
    let num_participants = sp1_zkvm::io::read::<u32>();
    let num_winners: u32 = sp1_zkvm::io::read::<u32>();

    // The drand network provides randomness outputs as a byte array with 32 bytes of data
    let randomness = sp1_zkvm::io::read::<[u8; 32]>();
    let randomness_as_u64 = u64::from_le_bytes(randomness[..8].try_into().unwrap());

    // Run the raffle and get winners
    let winners = raffle(num_participants, num_winners, randomness_as_u64);

    // Calculate Merkle root for winners
    let merkle_root = calculate_merkle_root(&winners);

    let pub_val = PubValStruct {
        num_participants,
        num_winners,
        randomness: alloy_sol_types::private::FixedBytes(randomness),
        merkle_root: alloy_sol_types::private::FixedBytes(merkle_root),
    };

    let bytes = PubValStruct::abi_encode(&pub_val);

    // Commit to the public values of the program
    sp1_zkvm::io::commit_slice(&bytes);
}

sol! {
    /// The public values encoded as a struct that can be easily deserialized inside Solidity.
    struct PubValStruct {
        uint32 num_participants;
        uint32 num_winners;
        bytes32 randomness;
        bytes32 merkle_root;
    }
}

fn calculate_merkle_root(winners: &[u32]) -> [u8; 32] {
    let mut hashes: Vec<[u8; 32]> = winners
        .iter()
        .map(|&w| {
            let mut hasher = Sha256_9_8::new();
            hasher.update(w.to_le_bytes());
            hasher.finalize().into()
        })
        .collect();

    while hashes.len() > 1 {
        let mut new_hashes = Vec::new();
        for chunk in hashes.chunks(2) {
            let mut hasher = Sha256_9_8::new();
            hasher.update(&chunk[0]);
            if chunk.len() > 1 {
                hasher.update(&chunk[1]);
            } else {
                // If there's no right child, hash with self
                hasher.update(&chunk[0]);
            }
            new_hashes.push(hasher.finalize().into());
        }
        hashes = new_hashes;
    }

    hashes[0]
}
