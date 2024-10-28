use alloy_sol_types::{sol, SolType};
use clap::Parser;
use sp1_sdk::{ProverClient, SP1ProofWithPublicValues, SP1Stdin};

/// The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
pub const ELF: &[u8] = include_bytes!("../../../elf/riscv32im-succinct-zkvm-elf");

/// The arguments for the command.
#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
struct Args {
    #[clap(long)]
    execute: bool,

    #[clap(long)]
    prove: bool,

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
    let args = Args::parse();

    // Setup the inputs.
    let mut stdin = SP1Stdin::new();
    stdin.write(&args.num_participants);
    stdin.write(&args.num_winners);
    stdin.write(&args.randomness);

    println!("Number of participants: {}", args.num_participants);
    println!("Number of winners: {}", args.num_winners);
    println!("Random seed: 0x{}", hex::encode(args.randomness));

    // Generate and verify the proof
    let client: ProverClient = ProverClient::new();

    if args.execute {
        // Execute the program
        let (output, report) = client.execute(ELF, stdin).run().unwrap();
        println!("Program executed successfully.");

        // Read the output.
        let decoded = PubValStruct::abi_decode(output.as_slice(), true).unwrap();
        let PubValStruct {
            num_participants,
            num_winners,
            randomness,
            merkle_root,
        } = decoded;

        // Verify that the output matches the input (remove random_seed check since it's now transformed)
        assert_eq!(
            num_participants, args.num_participants,
            "Number of participants mismatch"
        );
        assert_eq!(randomness, args.randomness, "Randomness mismatch");
        assert_eq!(num_winners, args.num_winners, "Number of winners mismatch");

        println!("Randomness: 0x{}", hex::encode(randomness));
        println!("Merkle Root: 0x{}", hex::encode(merkle_root));

        // Record the number of cycles executed.
        println!("Number of cycles: {}", report.total_instruction_count());
    } else if args.prove {
        // Setup the program for proving.
        let (pk, vk) = client.setup(ELF);

        let proof: SP1ProofWithPublicValues = client.prove(&pk, stdin).run().unwrap();

        client.verify(&proof, &vk).expect("verification failed");

        // Test a round trip of proof serialization and deserialization.
        proof
            .save("per-raffle-proof.bin")
            .expect("saving proof failed");
        let deserialized_proof =
            SP1ProofWithPublicValues::load("per-raffle-proof.bin").expect("loading proof failed");

        // Verify the deserialized proof.
        client
            .verify(&deserialized_proof, &vk)
            .expect("verification failed");

        println!("Successfully generated and verified proof for the raffle!");
    } else {
        println!("Please specify either --execute or --prove");
    }
}
