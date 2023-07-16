build:
	go build -o bin/main main.go
run:
	go run main.go
clean:
	rm -rf bin/main
debug:
	dlv debug --api-version=2 --log=true --config=debug.json
tidy:
	go mod tidy
proto:
	rm -rf pb/*
	protoc --proto_path=protos --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	protos/*.proto

.PHONY: debug proto
