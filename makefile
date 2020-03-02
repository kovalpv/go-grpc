proto:
	protoc -I ./api/proto/ --go_out=plugins=grpc:./pkg/api ./api/proto/transmitter.proto

server:
	go run ./cmd/server/main.go

client:
	go run ./cmd/client/main.go 0.0.0.0:8081