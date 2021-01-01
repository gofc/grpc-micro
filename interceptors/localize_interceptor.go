package interceptors

import (
	"context"
	"github.com/gofc/grpc-micro/code"
	"github.com/gofc/grpc-micro/locale"
	"github.com/gofc/grpc-micro/localizer"
	"github.com/gofc/grpc-micro/logger"
	pbbase "github.com/gofc/grpc-micro/proto"
	"github.com/gofc/grpc-micro/validations"
	"google.golang.org/grpc/metadata"
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//LocalizeInterceptor エラーメッセージを通訳するinterceptor
func LocalizeInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			err = localizeError(ctx, err)
		}
		return
	}
}

func getTimezone(ctx context.Context) string {
	return valueFromContext(ctx, "timezone", "UTC")
}

func getAcceptLanguage(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return locale.AcceptLanguage.Default
	}
	acceptLanguage := md.Get("accept-language")
	if len(acceptLanguage) == 0 {
		return locale.AcceptLanguage.Default
	}
	return locale.AcceptLanguage.ExtractLocale(acceptLanguage)
}

func valueFromContext(ctx context.Context, key string, def string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return def
	}
	v, ok := md[key]
	if !ok || len(v) == 0 {
		return def
	}
	return v[0]
}

func localizeError(ctx context.Context, err error) error {
	// -- エラーをローカライズしながら詰め直す

	st := status.Convert(err)
	if st.Code() == codes.OK {
		return nil
	}
	defSt := status.New(st.Code(), st.Code().String())

	// エラー詳細が無かった場合は詳細無しのエラーを返却する
	if len(st.Details()) == 0 {
		return code.NewGRPCWithStatus(defSt, err)
	}

	// エラー詳細をローカライズ
	lo := localizer.New(getTimezone(ctx), getAcceptLanguage(ctx))
	details, msg, err := localizeDetail(st.Details(), lo)
	if err != nil {
		// エラー詳細の詰め直しが失敗した場合は詳細無しのエラーを返却する
		logger.Error(ctx, "Error Details Adding failed.", zap.Error(err))
		return code.NewGRPCWithStatus(defSt, err)
	}

	newSt := status.New(st.Code(), msg)
	stWithDetail, err := newSt.WithDetails(details...)
	if err != nil {
		// エラー詳細の詰め直しが失敗した場合は詳細無しのエラーを返却する
		logger.Error(ctx, "Error Details Adding failed.", zap.Error(err))
		return code.NewGRPCWithStatus(newSt, err)
	}

	return code.NewGRPCWithStatus(stWithDetail, err)
}

func localizeDetail(details []interface{}, localizer localizer.Localizer) (v []proto.Message, message string, err error) {
	v = make([]proto.Message, 0)
	msgs := make([]string, 0)
	for _, detail := range details {
		switch d := detail.(type) {
		case *pbbase.ValidationError:
			p, msg := localizeValidationError(d, localizer)
			v = append(v, p)
			msgs = append(msgs, msg)
		case *pbbase.ApplicationError:
			p, msg := localizeAppError(d, localizer)
			v = append(v, p)
			msgs = append(msgs, msg)
		default:
			err = errors.Errorf("unexpected detail. %v", detail)
			return
		}
	}
	message = strings.Join(msgs, "\n")
	return
}

func localizeAppError(v *pbbase.ApplicationError, localizer localizer.Localizer) (proto.Message, string) {
	e, err := code.FromMessage(v)
	if err != nil {
		v.Message = code.Code(v.Code).String()
		return v, v.Message
	}
	v.Message = e.Localize(localizer)
	return v, v.Message
}

func localizeValidationError(v *pbbase.ValidationError, lo localizer.Localizer) (proto.Message, string) {
	ps, err := code.AnyToErrorParams(v.Params)
	if err != nil {
		v.Message = code.Code(v.Code).LocalizeKey()
	}

	details := validations.Detail{
		ErrorCode: code.Code(v.Code),
		Params:    ps,
	}
	v.Message = details.Localize(lo)

	return v, v.Message
}
