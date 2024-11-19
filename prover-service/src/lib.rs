pub mod proto {
    // Include the generated proto code
    pub mod prover {
        tonic::include_proto!("prover");
    }
}

pub mod service;

// Re-export common types
pub use proto::prover::{ProofSystem, ProveRequest, ProveResponse};
