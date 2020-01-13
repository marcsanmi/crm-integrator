build:
	protoc --proto_path=proto --go_out=plugins=grpc:proto customer.proto