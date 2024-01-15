UNAME := $(shell uname)

ifeq ($(UNAME), Linux)
GRPC_PROTOC_PLUGIN=${HOME}/.nuget/packages/grpc.tools/2.60.0/tools/linux_x64
else ifeq ($(UNAME), Darwin)
GRPC_PROTOC_PLUGIN=${HOME}/.nuget/packages/grpc.tools/2.60.0/tools/macosx_x64
else
GRPC_PROTOC_PLUGIN=${HOME}/.nuget/packages/grpc.tools/2.60.0/tools/windows_x64
endif

gen:
	protoc --csharp_out=. --grpc_out=. --plugin=protoc-gen-grpc=${GRPC_PROTOC_PLUGIN}/grpc_csharp_plugin reverse.proto

clean:
	dotnet clean