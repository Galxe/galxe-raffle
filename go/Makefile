PROTO_DIR = ../prover-service/proto
OUT_DIR = pkg/prover

.PHONY: all proto clean

all: proto

# Generate protobuf code
proto:
	protoc --go_out=$(OUT_DIR) \
				-I=$(PROTO_DIR) \
                --go_opt=paths=source_relative \
                --go-grpc_out=$(OUT_DIR) \
                --go-grpc_opt=paths=source_relative \
                $(PROTO_DIR)/*.proto

# Clean generated protobuf code
clean:
	rm -rf $(OUT_DIR)/*.pb.go
