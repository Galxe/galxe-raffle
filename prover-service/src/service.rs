use alloy_sol_types::sol;
use alloy_sol_types::SolType;
use prometheus::{register_histogram_vec, register_int_counter_vec, HistogramVec, IntCounterVec};
use sp1_sdk::network::prover::NetworkProver;
use sp1_sdk::network::FulfillmentStrategy;
use sp1_sdk::Prover;
use sp1_sdk::SP1ProofMode;
use sp1_sdk::SP1ProvingKey;
use sp1_sdk::SP1Stdin;

use std::time::Duration;
use tonic::{Request, Response, Status};

use crate::proto::prover::prover_service_server::ProverService;
use crate::proto::prover::{ProofSystem, ProveRequest, ProveResponse};
use crate::utils::hex::parse_hex_string;

const METRIC_LABEL_SUCCESS: &str = "success";
const METRIC_LABEL_ERROR: &str = "error";

pub struct ProverServiceImpl {
    pk: SP1ProvingKey,
    prover: NetworkProver,
    timeout_secs: u64,
}

impl ProverServiceImpl {
    pub fn new(
        private_key: &str,
        rpc_url: &str,
        elf_path: &str,
        timeout_secs: u64,
    ) -> std::io::Result<Self> {
        let prover = NetworkProver::new(private_key, rpc_url);
        let elf = std::fs::read(elf_path)?;
        let (pk, _vk) = prover.setup(elf.as_slice());
        Ok(Self {
            pk,
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

// Add metrics as lazy static
lazy_static::lazy_static! {
    pub static ref PROVE_REQUEST_DURATION: HistogramVec = register_histogram_vec!(
        "prove_request_duration_seconds",
        "Time taken to process prove requests",
        &["status"]
    ).unwrap();

    pub static ref PROVE_REQUEST_COUNTER: IntCounterVec = register_int_counter_vec!(
        "prove_request_total",
        "Total number of prove requests",
        &["status"]
    ).unwrap();
}

#[tonic::async_trait]
impl ProverService for ProverServiceImpl {
    async fn prove(
        &self,
        request: Request<ProveRequest>,
    ) -> Result<Response<ProveResponse>, Status> {
        // Start timing the request
        let start_time = std::time::Instant::now();

        let result: Result<Response<ProveResponse>, Status> = async {
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
                ProofSystem::Plonk => SP1ProofMode::Plonk,
                ProofSystem::Groth16 => SP1ProofMode::Groth16,
            };

            let proof_id = self
                .prover
                .prove(&self.pk, &stdin)
                .strategy(FulfillmentStrategy::Auction)
                .skip_simulation(false)
                .cycle_limit(140_000)
                .gas_limit(3_000_000)
                .mode(proof_mode)
                .request()
                .map_err(|e| Status::internal(e.to_string()))?;

            log::info!(
                "Proof requested. ID: {}, System: {}",
                proof_id,
                proof_system.as_str_name()
            );

            // Wait for the proof with timeout
            let proof = self
                .prover
                .wait_proof(proof_id, Some(Duration::from_secs(self.timeout_secs)), Some(Duration::from_secs(self.timeout_secs)))
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

            let balance = match self.prover.get_balance().await {
                Ok(balance) => balance,
                Err(e) => {
                    log::error!("Failed to get balance: {}", e);
                    Default::default() // Returns 0 for numeric types
                }
            };

            Ok(Response::new(ProveResponse {
                proof_id: proof_id.to_string(),
                public_values: format!("0x{}", hex::encode(bytes)),
                proof: format!("0x{}", hex::encode(proof.bytes())),
                num_participants,
                num_winners,
                randomness: format!("0x{}", hex::encode(&randomness.as_slice())),
                merkle_root: format!("0x{}", hex::encode(merkle_root.as_slice())),
                balance: balance.to_string(),
            }))
        }
        .await;

        // Create timer only after we know the result
        let request_duration = start_time.elapsed().as_secs_f64();
        let label = if result.is_ok() {
            METRIC_LABEL_SUCCESS
        } else {
            METRIC_LABEL_ERROR
        };
        PROVE_REQUEST_DURATION
            .with_label_values(&[label])
            .observe(request_duration);
        PROVE_REQUEST_COUNTER.with_label_values(&[label]).inc();
        result
    }
}
