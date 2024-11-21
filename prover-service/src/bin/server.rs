use clap::Parser;
use prover_service::proto::prover::prover_service_server::ProverServiceServer;
use prover_service::service::ProverServiceImpl;
use tonic::transport::Server;
use tonic_health::server::health_reporter;

#[derive(Parser)]
#[command(
    author = "Galxe",
    version = "0.0.1",
    about = "GRPC prover server for generating zero-knowledge SP1 proofs using the Prover Network",
    long_about = None
)]
struct Args {
    /// Port number for the server
    #[arg(long, default_value = "9090", env = "PORT")]
    port: u16,

    /// Private key of the prover (must be allowlisted by the prover network for now)
    /// Can be provided via PRIVATE_KEY environment variable
    #[arg(long, env = "PRIVATE_KEY")]
    private_key: String,

    /// Path to the ELF file
    #[arg(
        long,
        default_value = "elf/riscv32im-succinct-zkvm-elf",
        env = "ELF_PATH"
    )]
    elf_path: String,

    /// Timeout for waiting for the proof in seconds
    #[arg(long, default_value = "300", env = "TIMEOUT_SECS")]
    timeout_secs: u64,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Set up logging
    env_logger::init_from_env(env_logger::Env::default().default_filter_or("info"));

    // Parse arguments
    let args = Args::parse();

    // Start the server
    let addr = format!("0.0.0.0:{}", args.port).parse()?;
    let service: ProverServiceImpl =
        ProverServiceImpl::new(&args.private_key, &args.elf_path, args.timeout_secs)?;

    // Set up health checks
    let (mut health_reporter, health_service) = health_reporter();
    health_reporter
        .set_serving::<ProverServiceServer<ProverServiceImpl>>()
        .await;

    // Serve requests
    log::info!("Health Server and GRPC Prover Server listening on {}", addr);
    Server::builder()
        .add_service(health_service)
        .add_service(ProverServiceServer::new(service))
        .serve(addr)
        .await?;
    Ok(())
}
