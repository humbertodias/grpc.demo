# https://github.com/protocolbuffers/protobuf/releases/

UNAME := $(shell uname)


PROTOC_VERSION=25.2
#PROTOC_ARCH=aarch_64
PROTOC_ARCH=x86_64

GO_VERSION=1.21.6
#GO_ARCH=arm64
GO_ARCH=amd64

PROTOC_URL=https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-osx-${PROTOC_ARCH}.zip
GO_URL=https://go.dev/dl/go${GO_VERSION}.darwin-${GO_ARCH}.tar.gz
ifeq ($(UNAME), Linux)
PROTOC_URL=https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-${PROTOC_ARCH}.zip
GO_URL=https://go.dev/dl/go${GO_VERSION}.linux-${GO_ARCH}.tar.gz
endif

get-go:
	curl -o go.tar.gz -L "${GO_URL}"
	rm -rf /usr/local/go && tar -C /usr/local -xzf go.tar.gz
	rm -f go.tar.gz

get-protoc:
	curl -o protoc.zip -L "${PROTOC_URL}"
	unzip protoc.zip -d ~/go
	rm -f protoc.zip

profile:
	echo 'export PATH=~/go/bin:\$PATH:/usr/local/go/bin' >> ~/.bashrc
	source ~/.bashrc

get-plugin:
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	dotnet tool install -g dotnet-grpc
	python3 -m pip install grpcio-tools

gen-server:
	cd go.server && make gen
	cd python.server && make gen

gen-client:
	cd csharp.client && make gen


get-grpcui:
	go install github.com/fullstorydev/grpcui/cmd/grpcui@latest