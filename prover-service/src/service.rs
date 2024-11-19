use alloy_sol_types::sol;
use alloy_sol_types::SolType;
use sp1_sdk::network::proto::network::ProofMode;
use sp1_sdk::network::prover::NetworkProver;
use sp1_sdk::SP1Stdin;
use std::time::Duration;
use tonic::{Request, Response, Status};

use crate::proto::prover::prover_service_server::ProverService;
use crate::proto::prover::{ProofSystem, ProveRequest, ProveResponse};

pub struct ProverServiceImpl {
    elf: Vec<u8>,
    prover: NetworkProver,
    timeout_secs: u64,
}

impl ProverServiceImpl {
    pub fn new(private_key: &str, elf_path: &str, timeout_secs: u64) -> std::io::Result<Self> {
        let prover = NetworkProver::new_from_key(private_key);
        let elf = std::fs::read(elf_path)?;
        Ok(Self {
            elf,
            prover,
            timeout_secs,
        })
    }
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

#[tonic::async_trait]
impl ProverService for ProverServiceImpl {
    async fn prove(
        &self,
        request: Request<ProveRequest>,
    ) -> Result<Response<ProveResponse>, Status> {
        let req = request.into_inner();

        let randomness = parse_hex_string(&req.randomness)
            .map_err(|e| Status::invalid_argument(&format!("Invalid hex randomness: {}", e)))?;

        let proof_system = ProofSystem::try_from(req.system)
            .map_err(|e| Status::invalid_argument(&format!("Invalid proof system: {}", e)))?;

        let mut stdin = SP1Stdin::new();
        stdin.write(&req.num_participants);
        stdin.write(&req.num_winners);
        stdin.write(&randomness);

        let proof_mode = match proof_system {
            ProofSystem::Unspecified => {
                return Err(Status::invalid_argument("Proof system must be specified"));
            }
            ProofSystem::Plonk => ProofMode::Plonk,
            ProofSystem::Groth16 => ProofMode::Groth16,
        };

        let proof_id = self
            .prover
            .request_proof(self.elf.as_slice(), stdin, proof_mode)
            .await
            .map_err(|e| Status::internal(e.to_string()))?;

        log::info!(
            "Proof requested. ID: {}, System: {}",
            proof_id,
            proof_system.as_str_name()
        );

        // Wait for the proof with timeout
        let proof = self
            .prover
            .wait_proof(&proof_id, Some(Duration::from_secs(self.timeout_secs)))
            .await
            .map_err(|e| Status::internal(e.to_string()))?;

        let bytes = proof.public_values.as_slice();
        let PubValStruct {
            num_participants,
            num_winners,
            randomness,
            merkle_root,
        } = PubValStruct::abi_decode(bytes, false).unwrap();

        log::info!(
            "{}: Num Participants: {}, Num Winners: {}, Randomness: 0x{}, Merkle Root: 0x{}",
            proof_id,
            num_participants,
            num_winners,
            hex::encode(randomness.as_slice()),
            hex::encode(merkle_root.as_slice())
        );

        Ok(Response::new(ProveResponse {
            proof_id,
            public_values: format!("0x{}", hex::encode(bytes)),
            proof: format!("0x{}", hex::encode(proof.bytes())),
            num_participants,
            num_winners,
            randomness: format!("0x{}", hex::encode(&randomness.as_slice())),
            merkle_root: format!("0x{}", hex::encode(merkle_root.as_slice())),
        }))
    }
}
