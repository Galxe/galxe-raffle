pub mod proto {
    // Include the generated proto code
    pub mod prover {
        tonic::include_proto!("prover");
    }
}

pub mod service;
pub mod utils;
// Re-export common types
pub use proto::prover::{ProofSystem, ProveRequest, ProveResponse};
