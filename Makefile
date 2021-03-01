.PHONY: proto
proto:
	protoc -I. \
        -I/home/vielen/GoProjects/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.2.0/third_party/googleapis \
		-I/home/vielen/GoProjects/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.2.0 \
        --openapiv2_out=disable_default_errors=true,allow_merge=true:. --go_out=:. --micro_out=components="micro|http":. proto/*.proto