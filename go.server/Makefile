gen:
	protoc --go_out=. --go-grpc_out=. reverse.proto

server:
	go run main.go

client:
	grpcui --proto reverse.proto -plaintext localhost:50051

