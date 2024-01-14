# https://github.com/protocolbuffers/protobuf/releases/

get-protoc:
	curl -o protoc.zip -L "https://github.com/protocolbuffers/protobuf/releases/download/v25.2/protoc-25.2-osx-aarch_64.zip"
	unzip protoc.zip -d ~/go
	rm -f protoc.zip

profile:
	echo 'export PATH=~/go/bin:$PATH' >> ~/.zshrc
	source ~/.zshrc
	echo 'export PATH=~/go/bin:$PATH' >> ~/.bashrc
	source ~/.bashrc

get-plugin:
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

gen:
	protoc \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/reverse.proto

gen-clean:
	rm proto/*.pb.*

