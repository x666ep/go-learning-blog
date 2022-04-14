all:
	mkdir -p "pkg"
	protoc -I. \
		-I./third_party/googleapis \
		--go_out ./pkg --go_opt paths=source_relative \
		--go-grpc_out ./pkg --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./pkg --grpc-gateway_opt paths=source_relative \
		--openapiv2_out ./pkg --openapiv2_opt logtostderr=true \
		./api/go-learning-blog/v1/api.proto
