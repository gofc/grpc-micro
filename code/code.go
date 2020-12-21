package code

import (
	"strconv"

	"github.com/gofc/grpc-micro/localizer"
	pbbase "github.com/gofc/grpc-micro/proto"

	"google.golang.org/grpc/codes"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

// Code エラーコード
type Code int

// Error Appエラーを作成する
func (i Code) Error() string {
	return i.String()
}

// AppError Appエラーを作成する
func (i Code) AppError(params ...interface{}) error {
	return Error(i, params...)
}

//GrpcError Grpcエラーを作成する
func (i Code) GrpcError(code codes.Code, params ...interface{}) error {
	return NewGRPC(code, Error(i, params...))
}

//LocalizeKey ローカライズ用のメッセージID
func (i Code) LocalizeKey() string {
	return i.String()
}

//String コードを文字列へ変更
func (i Code) String() string {
	return ""
}

// AppError システム全体で使うつもり.
type AppError struct {
	error
	Code   Code
	Params []ErrorParam
}

// Error エラーコードとパラメータを指定して作成
// paramsに指定された値は、基本的にfmt.Sprintf()されて文字列、またはローカライズ可能のWordとして設定される.
func Error(code Code, params ...interface{}) error {
	return newError(code, ToErrorParams(params))
}

//WrapError runtimeエラー、エラーコードとパラメータを指定して作成
func WrapError(err error, code Code, params ...interface{}) error {
	return &AppError{
		error:  errors.Wrapf(err, "ApplicationError: { code: %v, param:%v }", code, params),
		Code:   code,
		Params: ToErrorParams(params),
	}
}

func newError(code Code, params []ErrorParam) *AppError {
	return &AppError{
		error:  errors.Errorf("ApplicationError: { code: %v, param:%v }", code, params),
		Code:   code,
		Params: params,
	}
}

// FromMessage メッセージから作成
func FromMessage(msg *pbbase.ApplicationError) (*AppError, error) {
	params, err := AnyToErrorParams(msg.Params)
	if err != nil {
		return nil, err
	}

	return newError(Code(msg.Code), params), nil
}

// ToGrpcMessage gRPCメッセージに変換
func (e *AppError) ToGrpcMessage() (proto.Message, error) {
	anyList, err := ToAny(e.Params)
	if err != nil {
		return nil, err
	}

	return &pbbase.ApplicationError{
		Code:   uint64(e.Code),
		Params: anyList,
	}, nil
}

// Localize ローカライズ
func (e *AppError) Localize(l localizer.Localizer) string {
	params := make(map[string]interface{}, len(e.Params))
	for i, p := range e.Params {
		key := "Param" + strconv.Itoa(i)
		params[key] = p
	}

	localized, err := l.Localize(e.Code.LocalizeKey(), params)
	if err != nil {
		return e.Code.LocalizeKey()
	}
	return localized
}
