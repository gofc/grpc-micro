GO_PATH=`go env GOPATH`

install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ${GO_PATH}/bin v1.25.0
	GO111MODULE=off go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	GO111MODULE=off go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	GO111MODULE=off go get -u github.com/golang/protobuf/protoc-gen-go
	GO111MODULE=off go get -u golang.org/x/tools/cmd/stringer

deps:
	go mod tidy

gen: gen-proto
	@echo "generate all done"

gen-proto:
	@protoc -I. -I${GO_PATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    		--go_out=plugins=grpc,paths=source_relative:. \
    		--swagger_out=logtostderr=true,json_names_for_fields=true:. \
    		--grpc-gateway_out=logtostderr=true,paths=source_relative:. ./proto/v1/pbcomm/*.proto
	@protoc -I. -I${GO_PATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--go_out=plugins=grpc,paths=source_relative:. \
			--swagger_out=logtostderr=true,json_names_for_fields=true:. \
			--grpc-gateway_out=logtostderr=true,paths=source_relative:. ./proto/v1/*.proto
	@protoc -I. -I${GO_PATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--go_out=plugins=grpc,paths=source_relative:. \
			--swagger_out=logtostderr=true,json_names_for_fields=true:. \
			--grpc-gateway_out=logtostderr=true,paths=source_relative:. ./proto/v1/pbadmin/*.proto
	@echo "proto generate done"

build-app:
	@mkdir -p build/bin
	$(call build-app-target,foo)
	$(call build-app-target,cli)
	$(call build-app-target,restgw)

build-app-specify:
	$(call build-app-target,$(name))

build-image:
	$(call build-docker-image,foo)
	$(call build-docker-image,restgw)

run-all: build-app run-docker
	@

run-docker:
	docker-compose down -v
	docker-compose up

define build-docker-image
	docker build -t gofc/images:grpc-micro-$(1)-latest -f docker-files/$(1).Dockerfile .
endef

define build-app-target
	@echo "build $(1) service"
	@GOOS=linux go build -i -o build/bin/$(1) cmd/$(1)/main.go
endef
