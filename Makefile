GO_PATH=`go env GOPATH`
current_dir = $(shell pwd)

install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GO_PATH}/bin v1.32.2

deps:
	go mod download

go-lint:
	golangci-lint run --fix

gen:
	docker run -it --rm -v ${current_dir}:/workspace goforcloud/go-micro-proto-generator:1.0.0 make gen-proto

gen-proto:
	$(call gen-proto-target,./proto/*.proto)

define gen-proto-target
	protoc -I/usr/local/include -I. \
		--go_out ./ --go_opt paths=source_relative $(1)
endef
