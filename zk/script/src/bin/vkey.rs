use sp1_sdk::blocking::{Prover, ProverClient};
use sp1_sdk::{include_elf, Elf, HashableKey, ProvingKey};

/// The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
pub const FIBONACCI_ELF: Elf = include_elf!("galxe-raffle");

fn main() {
    let prover = ProverClient::builder().cpu().build();
    let pk = prover.setup(FIBONACCI_ELF).expect("failed to setup prover");
    println!("{}", pk.verifying_key().bytes32());
}
