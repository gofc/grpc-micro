deps:
	go mod tidy

gen: gen-proto
	@echo "generate all done"

gen-proto:
	@protoc --go_out=plugins=grpc,paths=source_relative:. ./proto/v1/pbcomm/*.proto
	@protoc --go_out=plugins=grpc,paths=source_relative:. ./proto/v1/*.proto
	@protoc --go_out=plugins=grpc,paths=source_relative:. ./proto/v1/pbadmin/*.proto
	@echo "proto generate done"
