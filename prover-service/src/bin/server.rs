use clap::Parser;
use prover_service::proto::prover::prover_service_server::ProverServiceServer;
use prover_service::service::ProverServiceImpl;
use tonic::transport::Server;

#[derive(Parser)]
#[command(
    author = "Galxe",
    version = "0.0.1",
    about = "GRPC prover server for generating zero-knowledge SP1 proofs using the Prover Network",
    long_about = None
)]
struct Args {
    /// Port number for the server
    #[arg(long, default_value = "9090")]
    port: u16,

    /// Private key of the prover (must be allowlisted by the prover network for now)
    #[arg(long)]
    private_key: String,

    /// Path to the ELF file
    #[arg(long, default_value = "elf/riscv32im-succinct-zkvm-elf")]
    elf_path: String,

    /// Timeout for waiting for the proof in seconds
    #[arg(long, default_value = "300")]
    timeout_secs: u64,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Set up logging
    env_logger::init_from_env(env_logger::Env::default().default_filter_or("info"));

    // Parse arguments
    let args = Args::parse();

    // Start the server
    let addr = format!("[::1]:{}", args.port).parse()?;
    let service: ProverServiceImpl =
        ProverServiceImpl::new(&args.private_key, &args.elf_path, args.timeout_secs)?;
    log::info!("ProverService listening on {}", addr);

    // Serve requests
    Server::builder()
        .add_service(ProverServiceServer::new(service))
        .serve(addr)
        .await?;
    Ok(())
}
