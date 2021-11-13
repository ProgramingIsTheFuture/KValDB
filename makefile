run:
	go run cmd/main.go

test:
	go test ./...

grpc:
	protoc --go_out=external/gRPC/. --go_opt=paths=source_relative \
    --go-grpc_out=external/gRPC/. --go-grpc_opt=paths=source_relative \
    proto/*.proto
