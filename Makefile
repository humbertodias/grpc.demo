# https://github.com/protocolbuffers/protobuf/releases/

get-protoc:
	curl -o protoc.zip -L "https://github.com/protocolbuffers/protobuf/releases/download/v25.2/protoc-25.2-osx-aarch_64.zip"
	unzip protoc.zip -d "tools"
	rm -f protoc.zip

get-plugin:
	GOBIN=`pwd`/tools/bin go install github.com/golang/protobuf/protoc-gen-go@latest

# gen:
# 	PATH=tools/bin:$$PATH protoc -Iproto proto/reverse.proto \
# 	--go_opt=Mproto/reverse.proto=github.com/humbertodias/grpc.demo/proto/reverse \
# 	--go_opt=paths=source_relative \
# 	--go_out=plugins=grpc:proto

# gen:
# 	PATH=tools/bin:$$PATH protoc \
# 	--go_opt=Mproto/reverse.proto=github.com/humbertodias/grpc.demo/proto/reverse \
# 	proto/reverse.proto \
# 	--go_out=plugins=grpc:proto

gen:
	PATH=tools/bin:$$PATH protoc --go_out=. proto/reverse.proto

gen-clean:
	rm proto/*.pb.*

