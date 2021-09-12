.PHONY: default
default:
	protoc \
    		-I=./ \
    		-I=../../api \
    		--pw-errors_out=./ \
    		--go_out=./ \
    		./example.proto