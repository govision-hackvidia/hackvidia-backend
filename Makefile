generate_mllm:
	@protoc \
		--proto_path=protobuf "protobuf/mllm.proto" \
		--go_out=go/protobuf/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=go/protobuf/ \
		--go-grpc_opt=paths=source_relative

	@python3 -m grpc_tools.protoc \
		--proto_path=protobuf "protobuf/mllm.proto" \
		--python_out=python/mllm \
		--grpc_python_out=python/mllm