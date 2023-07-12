build:
	go build main.go
run:
	go run main.go
clean:
	go run clean

debug:
	dlv debug --api-version=2 --log=true --config=debug.json

.PHONY: debug

proto:
	protoc --go_out=. --go-grpc_out=. api/proto/inventory/*.proto
