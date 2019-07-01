compile_protos:
	protoc -I . --php_out=php_client \
	--grpc_out=proto \
	--plugin=protoc-gen-grpc=bin/grpc_php_plugin proto/*.proto

build:
	go build -o server/server server/main.go