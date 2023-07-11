build:
	go build main.go
run:
	go run main.go
clean:
	go run clean
proto:
	protoc --go_out=. --go-grpc_out=. api/proto/inventory/*.proto
