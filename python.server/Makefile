gen:
	python3 -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. greeter.proto

server:
	python3 server.py &

client:
	grpcui --proto greeter.proto -plaintext localhost:50051