use clap::Parser;
use prover_service::proto::prover::prover_service_server::ProverServiceServer;
use prover_service::service::ProverServiceImpl;
use std::net::SocketAddr;
use tonic::codec::CompressionEncoding;
use tonic::transport::Server;
use tonic_health::server::health_reporter;
use warp::Filter;

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

    /// Port number for Prometheus metrics
    #[arg(long, default_value = "4014", env = "METRICS_PORT")]
    metrics_port: u16,
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

    // Start Prometheus metrics endpoint
    let metrics_addr = format!("0.0.0.0:{}", args.metrics_port);
    let metrics_route = warp::path!("metrics").and_then(|| async {
        use prometheus::Encoder;
        let encoder = prometheus::TextEncoder::new();
        let mut buffer = Vec::new();
        encoder.encode(&prometheus::gather(), &mut buffer).unwrap();
        Ok::<_, warp::Rejection>(String::from_utf8(buffer).unwrap())
    });

    tokio::spawn(async move {
        log::info!("Prometheus metrics server listening on {}", metrics_addr);
        warp::serve(metrics_route)
            .run(metrics_addr.parse::<SocketAddr>().unwrap())
            .await;
    });

    // Start prover service
    let prover_service = ProverServiceServer::new(service)
        .accept_compressed(CompressionEncoding::Gzip)
        .send_compressed(CompressionEncoding::Gzip);

    // Serve requests
    log::info!("Health Server and GRPC Prover Server listening on {}", addr);
    Server::builder()
        .add_service(health_service)
        .add_service(prover_service)
        .serve(addr)
        .await?;
    Ok(())
}
