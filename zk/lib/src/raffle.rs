use sha2_v0_9_8::{Digest as Digest_sha2_9_8, Sha256 as Sha256_9_8};
// use sha2_v0_10_8::{Digest as Digest_sha2_10_8, Sha256 as Sha256_10_8};
// use sha3_v0_10_8::{Digest as Digest_sha3_10_8, Keccak256 as Keccak256_10_8};
// use tiny_keccak::{Hasher, Keccak as Tiny_Keccak};
use std::collections::HashSet;

/// Naive raffle implementation
pub fn raffle_naive(num_participants: u32, num_winners: u32, random_seed: u64) -> Vec<u32> {
    let n = num_participants as usize;
    let m = num_winners as usize;

    let (select_count, invert) = if m > n / 2 { (n - m, true) } else { (m, false) };

    let mut selected = Vec::with_capacity(select_count);
    let mut used_numbers: HashSet<u32> = HashSet::with_capacity(select_count);
    let mut seed = random_seed;

    while selected.len() < select_count {
        let new_number = (seed % n as u64) as u32;
        if used_numbers.insert(new_number) {
            selected.push(new_number);
        }
        seed = hash_sha2(seed, new_number as u64);
    }

    if invert {
        let mut result = Vec::with_capacity(n - select_count);
        for i in 0..num_participants {
            if !used_numbers.contains(&i) {
                result.push(i);
            }
        }
        result
    } else {
        selected
    }
}

/// Raffle implementation using a modified Fisher-Yates shuffle
pub fn raffle_fisher_yates(num_participants: u32, num_winners: u32, random_seed: u64) -> Vec<u32> {
    let n = num_participants as usize;
    let m = num_winners as usize;

    let (shuffle_count, invert) = if m > n / 2 { (n - m, true) } else { (m, false) };

    let mut participants: Vec<u32> = (0..num_participants).collect();
    let mut seed = random_seed;

    for i in 0..shuffle_count {
        seed = hash_sha2(seed, i as u64);
        let j: usize = i + (seed % (n - i) as u64) as usize;
        participants.swap(i, j);
    }

    if invert {
        participants.into_iter().skip(shuffle_count).collect()
    } else {
        participants.into_iter().take(shuffle_count).collect()
    }
}

fn hash_sha2(seed: u64, value: u64) -> u64 {
    let mut sha256: Sha256_9_8 = Sha256_9_8::new();
    sha256.update(seed.to_le_bytes());
    sha256.update(value.to_le_bytes());
    let result = sha256.finalize();
    u64::from_le_bytes(result[..8].try_into().unwrap())
}

// fn hash_sha3(seed: u64, value: u64) -> u64 {
//     let mut keccak256: Keccak256_10_8 = Keccak256_10_8::new();
//     keccak256.update(seed.to_le_bytes());
//     keccak256.update(value.to_le_bytes());
//     let result = keccak256.finalize();
//     u64::from_le_bytes(result[..8].try_into().unwrap())
// }

// fn hash_tiny_keccak(seed: u64, value: u64) -> u64 {
//     let mut keccak = Tiny_Keccak::v256();
//     keccak.update(&seed.to_le_bytes());
//     keccak.update(&value.to_le_bytes());
//     let mut output = [0u8; 32];
//     keccak.finalize(&mut output);
//     u64::from_le_bytes(output[..8].try_into().unwrap())
// }
