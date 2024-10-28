use sha2_v0_9_8::{Digest as Digest_sha2_9_8, Sha256 as Sha256_9_8};
use std::collections::HashSet;

pub fn raffle(num_participants: u32, num_winners: u32, randomness: u64) -> Vec<u32> {
    let n = num_participants as usize;
    let m = num_winners as usize;

    let (select_count, invert) = if m > n / 2 { (n - m, true) } else { (m, false) };

    let mut selected = Vec::with_capacity(select_count);
    let mut used_numbers: HashSet<u32> = HashSet::with_capacity(select_count);
    let mut seed = randomness;

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

fn hash_sha2(seed: u64, value: u64) -> u64 {
    let mut sha256: Sha256_9_8 = Sha256_9_8::new();
    sha256.update(seed.to_le_bytes());
    sha256.update(value.to_le_bytes());
    let result = sha256.finalize();
    u64::from_le_bytes(result[..8].try_into().unwrap())
}
