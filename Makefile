# https://github.com/protocolbuffers/protobuf/releases/

UNAME := $(shell uname)


PROTOC_VERSION=25.2
PROTOC_ARCH=aarch_64
#PROTOC_ARCH=x86_64

GO_VERSION=1.21.6
#GO_ARCH=arm64
#GO_ARCH=amd64
GO_ARCH=arm64

PROTOC_URL=https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-win64.zip
GO_URL=https://go.dev/dl/go${GO_VERSION}.windows-${GO_ARCH}.zip

ifeq ($(UNAME), Linux)
PROTOC_URL=https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-${PROTOC_ARCH}.zip
GO_URL=https://go.dev/dl/go${GO_VERSION}.linux-${GO_ARCH}.tar.gz
else ifeq ($(UNAME), Darwin)
PROTOC_URL=https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-osx-${PROTOC_ARCH}.zip
GO_URL=https://go.dev/dl/go${GO_VERSION}.darwin-${GO_ARCH}.tar.gz
endif

get-go:
	curl -o go.tar.gz -L "${GO_URL}"
	rm -rf /usr/local/go && tar -C /usr/local -xzf go.tar.gz
	rm -f go.tar.gz

get-go-win:
	curl -o go.zip -L "${GO_URL}"
	rm -rf /usr/local/go && unzip -d /usr/local go.zip
	rm -f go.zip
	
get-protoc:
	curl -o protoc.zip -L "${PROTOC_URL}"
	unzip protoc.zip -d ~/go
	rm -f protoc.zip

profile:
	echo 'export PATH=~/go/bin:/usr/local/go/bin:$$PATH' >> ~/.bashrc
	source ${HOME}/.bashrc

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