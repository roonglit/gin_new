server:
	go run main.go

proto:
	rm -f grpc/pb/*.go
	protoc --proto_path=grpc/proto --go_out=grpc/pb --go_opt=paths=source_relative \
    --go-grpc_out=grpc/pb --go-grpc_opt=paths=source_relative \
    grpc/proto/*.proto

.PHONY: server proto