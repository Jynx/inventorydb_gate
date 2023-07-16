build:
	go build main.go
run:
	go run main.go
clean:
	go run clean
debug:
	dlv debug --api-version=2 --log=true --config=debug.json
proto:
	rm -rf pb/*
	protoc --proto_path=protobufs --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	protobufs/*.proto

.PHONY: debug proto
