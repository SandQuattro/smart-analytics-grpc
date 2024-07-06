generate:
	protoc --go_out=. --go-grpc_out=. proto/assistants/assistants.proto proto/users/userinfo.proto
.PHONY: generate