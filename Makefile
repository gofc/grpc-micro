GO_PATH=`go env GOPATH`
available_gateways=" "
available_apps=" foo "
current_dir = $(shell pwd)

deps:
	go mod tidy

gen: gen-proto
	@echo "generate all done"

gen-proto:
	@protoc --go_out=plugins=grpc,paths=source_relative:. ./proto/v1/pbcomm/*.proto
	@protoc --go_out=plugins=grpc,paths=source_relative:. ./proto/v1/*.proto
	@protoc --go_out=plugins=grpc,paths=source_relative:. ./proto/v1/pbadmin/*.proto
	@echo "proto generate done"

build-app:
	@mkdir -p build/bin
	@./scripts/build-all.sh

build-app-specify:
ifeq ($(filter $(name),$(available_apps)),$(name))
	@GOOS=linux go build -i -o build/bin/${name} cmd/${name}/main.go
	@echo "service $(name) build success"
else ifeq ($(filter $(name),$(available_gateways)),$(name))
	@GOOS=linux go build -i -o build/bin/${name} cmd/${name}/main.go
	@GOOS=linux go build -i -o build/bin/${name}_proxy cmd/${name}/proxy/proxy.go
	@echo "gateway service $(name) build success"
else
	@echo "invalid app name"
endif

build-image:
	docker build -t jmz331/global-images:grpc-micro-foo-latest -f docker-files/foo.Dockerfile .

run-all: build-app
	docker-compose down -v
	docker-compose up
