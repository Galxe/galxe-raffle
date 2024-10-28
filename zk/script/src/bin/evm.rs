//! An end-to-end example of using the SP1 SDK to generate a proof of a program that can have an
//! EVM-Compatible proof generated which can be verified on-chain.
//!
//! You can run this script using the following command:
//! ```shell
//! RUST_LOG=info cargo run --release --bin evm -- --system groth16
//! ```
//! or
//! ```shell
//! RUST_LOG=info cargo run --release --bin evm -- --system plonk
//! ```

use alloy_sol_types::sol;
use alloy_sol_types::SolType;

use clap::{Parser, ValueEnum};
use serde::{Deserialize, Serialize};
use sp1_sdk::{HashableKey, ProverClient, SP1ProofWithPublicValues, SP1Stdin, SP1VerifyingKey};
use std::path::PathBuf;

/// The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
pub const ELF: &[u8] = include_bytes!("../../../elf/riscv32im-succinct-zkvm-elf");

/// The arguments for the EVM command.
#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
struct EVMArgs {
    #[clap(long, value_enum, default_value = "groth16")]
    system: ProofSystem,

    #[clap(long, default_value = "100")]
    num_participants: u32,

    #[clap(long, default_value = "10")]
    num_winners: u32,

    #[clap(
        long,
        default_value = "0x82e8b6bbf24681c9d3c928f988aa6eef88f41f164e5df290e3dca3b8f6ce3f07",
        value_parser = parse_hex_string
    )]
    randomness: [u8; 32],
}

/// Enum representing the available proof systems
#[derive(Copy, Clone, PartialEq, Eq, PartialOrd, Ord, ValueEnum, Debug)]
enum ProofSystem {
    Plonk,
    Groth16,
}

/// A fixture that can be used to test the verification of SP1 zkVM proofs inside Solidity.
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
struct SP1ProofFixture {
    num_participants: u32,
    num_winners: u32,
    randomness: String,
    merkle_root: String,
    vkey: String,
    public_values: String,
    proof: String,
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

/// Parse a hex string into a 32-byte array
fn parse_hex_string(s: &str) -> Result<[u8; 32], String> {
    let s = s.strip_prefix("0x").unwrap_or(s);
    if s.len() != 64 {
        return Err("Randomness must be a 32-byte (64 character) hex string".to_string());
    }

    let bytes = hex::decode(s).map_err(|e| format!("Failed to decode hex string: {}", e))?;

    bytes
        .try_into()
        .map_err(|_| "Failed to convert to 32 byte array".to_string())
}

fn main() {
    // Setup the logger.
    sp1_sdk::utils::setup_logger();

    // Parse the command line arguments.
    let args = EVMArgs::parse();

    // Setup the prover client.
    let client = ProverClient::new();

    // Setup the program.
    let (pk, vk) = client.setup(ELF);

    // Setup the inputs.
    let mut stdin = SP1Stdin::new();
    stdin.write(&args.num_participants);
    stdin.write(&args.num_winners);
    stdin.write(&args.randomness);

    println!("Num Participants: {}", args.num_participants);
    println!("Num Winners: {}", args.num_winners);
    println!("Randomness: 0x{}", hex::encode(args.randomness));
    println!("Proof System: {:?}", args.system);

    // Generate the proof based on the selected proof system.
    let proof = match args.system {
        ProofSystem::Plonk => client.prove(&pk, stdin).plonk().run(),
        ProofSystem::Groth16 => client.prove(&pk, stdin).groth16().run(),
    }
    .expect("failed to generate proof");

    create_proof_fixture(&proof, &vk, args.system);
}

/// Create a fixture for the given proof.
fn create_proof_fixture(
    proof: &SP1ProofWithPublicValues,
    vk: &SP1VerifyingKey,
    system: ProofSystem,
) {
    // Deserialize the public values.
    let bytes = proof.public_values.as_slice();
    let PubValStruct {
        num_participants,
        num_winners,
        randomness,
        merkle_root,
    } = PubValStruct::abi_decode(bytes, false).unwrap();

    // Create the testing fixture so we can test things end-to-end.
    let fixture = SP1ProofFixture {
        num_participants,
        num_winners,
        randomness: format!("0x{}", hex::encode(randomness.as_slice())),
        merkle_root: format!("0x{}", hex::encode(merkle_root.as_slice())),
        vkey: vk.bytes32().to_string(),
        public_values: format!("0x{}", hex::encode(bytes)),
        proof: format!("0x{}", hex::encode(proof.bytes())),
    };

    // The verification key is used to verify that the proof corresponds to the execution of the
    // program on the given input.
    //
    // Note that the verification key stays the same regardless of the input.
    println!("Verification Key: {}", fixture.vkey);

    // The public values are the values which are publicly committed to by the zkVM.
    //
    // If you need to expose the inputs or outputs of your program, you should commit them in
    // the public values.
    println!("Public Values: {}", fixture.public_values);

    // The proof proves to the verifier that the program was executed with some inputs that led to
    // the give public values.
    println!("Proof Bytes: {}", fixture.proof);

    // Save the fixture to a file.
    let fixture_path = PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("../fixtures");
    std::fs::create_dir_all(&fixture_path).expect("failed to create fixture path");
    std::fs::write(
        fixture_path.join(
            format!(
                "{:?}-{}-{}-{}-fixture.json",
                system,
                num_participants,
                num_winners,
                hex::encode(&randomness.as_slice())
            )
            .to_lowercase(),
        ),
        serde_json::to_string_pretty(&fixture).unwrap(),
    )
    .expect("failed to write fixture");
}
