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
	dotnet tool install -g dotnet-grpc

gen-server:
	cd go.server && make gen
	cd python.server && make gen

gen-client:
	cd csharp.client && make gen
