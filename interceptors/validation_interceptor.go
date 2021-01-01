package interceptors

import (
	"context"
	"github.com/gofc/grpc-micro/validations"

	"google.golang.org/grpc"
)

type validator interface {
	Validate() *validations.Result
}

//ValidationInterceptor requestをvalidateするinterceptor
func ValidationInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if r, ok := req.(validator); ok {
			validateResult := r.Validate()
			if !validateResult.IsOK() {
				return nil, validateResult.Error()
			}
		}
		return handler(ctx, req)
	}
}
