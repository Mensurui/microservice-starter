.PHONY: protos

protos:
	protoc -I protos/ protos/currency.proto \
    --go_out=paths=source_relative:protos/currency \
    --go-grpc_out=paths=source_relative:protos/currency
