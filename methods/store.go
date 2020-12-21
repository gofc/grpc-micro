package methods

import (
	"context"
	"fmt"

	pbbase "github.com/gofc/grpc-micro/proto"

	"github.com/gofc/grpc-micro/logger"
	"github.com/jhump/protoreflect/grpcreflect"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"
)

var methodStore map[string]*pbbase.OptionBase

//Init RPC service定義からOptionBaseメソッド設定を抽出して初期化する
func Init(ctx context.Context, srv *grpc.Server) {
	sds, err := grpcreflect.LoadServiceDescriptors(srv)
	if err != nil {
		logger.Error(ctx, "failed to LoadServiceDescriptors from grpc server", zap.Error(err))
		panic(err)
	}

	methodStore = make(map[string]*pbbase.OptionBase)
	for _, sd := range sds {
		if sd == nil {
			logger.Error(context.Background(), "methods init failed")
			continue
		}
		for _, md := range sd.GetMethods() {
			opts := md.GetMethodOptions()

			val, err := proto.GetExtension(opts, pbbase.E_OptionBase)

			if err == nil {
				option, ok := val.(*pbbase.OptionBase)
				if ok {
					methodStore[fmt.Sprintf("/%s/%s", sd.GetFullyQualifiedName(), md.GetName())] = option
				}
			}
		}
	}
	methodStore["/grpc.health.v1.Health/Check"] = &pbbase.OptionBase{
		IgnoreTokenVerify:   true,
		IgnoreLoggerPayload: true,
	}
}

//IsIgnoreTokenVerify 指定メッソドがToken認証無視する
func IsIgnoreTokenVerify(methodName string) bool {
	val, ok := methodStore[methodName]
	if ok {
		return val.IgnoreTokenVerify
	}
	return false
}

//IsIgnoreLoggerPayload 指定メッソドのPayload内容をログに出力しない
func IsIgnoreLoggerPayload(methodName string) bool {
	val, ok := methodStore[methodName]
	if ok {
		return val.IgnoreLoggerPayload
	}
	return false
}
