.PHONY: api
api: validator
	protoc \
    		--go_opt=module=github.com/joesonw/proto-web \
    		--proto_path=./api \
    		--go_out=./ \
    		--pw-validator_opt=module=github.com/joesonw/proto-web \
    		--pw-validator_out=./ \
    		openapi.proto

.PHONY: errors
errors: api
	protoc \
    		--go_opt=module=github.com/joesonw/proto-web \
    		--proto_path=./api \
    		--go_out=./ \
    		errors.proto
	go install ./cmd/protoc-gen-pw-errors


.PHONY: http-server
http-server: api
	go install ./cmd/protoc-gen-pw-http-server

.PHONY: http-server
openapi: api
	go install ./cmd/protoc-gen-pw-openapi

.PHONY: validator
validator:
	protoc \
    		--go_opt=module=github.com/joesonw/proto-web \
    		--proto_path=./api \
    		--go_out=./ \
    		validator.proto
	go install ./cmd/protoc-gen-pw-validator

