# Build stage
FROM rust:1-slim as builder

WORKDIR /usr/src/app

RUN apt-get update && apt-get install -y \
    libssl3 \
    ca-certificates \
    pkg-config \
    libssl-dev \
    protobuf-compiler \
    g++

COPY . .

RUN cargo build --release

# Runtime stage
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y libssl3 ca-certificates
COPY --from=builder /usr/src/app/target/release/server /usr/local/bin/server
COPY --from=builder /usr/src/app/elf/riscv32im-succinct-zkvm-elf /usr/local/bin/elf/

EXPOSE 9090 4014

CMD ["server", "--elf-path", "/usr/local/bin/elf/riscv32im-succinct-zkvm-elf"]
