.PHONY: default
default:
	protoc \
    		-I=./ \
    		-I=../../api \
    		--go-grpc_out=./ \
    		--go_out=./ \
    		--pw-http-server_out=./ \
    		--pw-openapi_out=./ \
    		./server.proto