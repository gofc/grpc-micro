package code

import (
	"fmt"
	"reflect"

	"github.com/gofc/grpc-micro/localizer"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"

	"github.com/golang/protobuf/proto"
)

// ErrorParam エラー情報のパラメータ.ローカライズ可能で、grpcで送信可能
type ErrorParam interface {
	localizer.Localizable

	// ToProtoMessage protoメッセージを作成
	ToProtoMessage() proto.Message

	// MessageType protoメッセージの型
	MessageType() reflect.Type

	// SetValue protoメッセージから値を詰める
	SetValue(msg proto.Message) error
}

func init() {
	registerErrorParam(&StringParam{})
	registerErrorParam(&LocalizeKeyParam{})
	registerErrorParam(&TimeParam{})
}

var messageToParamMap = map[reflect.Type]reflect.Type{}

// RegisterErrorParam 利用するパラメータを変換用のマップに登録しておく
func registerErrorParam(p ErrorParam) {
	messageToParamMap[p.MessageType()] = reflect.TypeOf(p)
}

// ToErrorParams ErrorParamsに変換
func ToErrorParams(ps []interface{}) []ErrorParam {
	l := len(ps)
	if l == 0 {
		return nil
	}
	eparams := make([]ErrorParam, len(ps))
	for i, p := range ps {
		eparams[i] = ToErrorParam(p)
	}
	return eparams
}

// ToErrorParam ErrorParamに変換
func ToErrorParam(p interface{}) ErrorParam {
	switch t := p.(type) {
	case ErrorParam:
		return t
	case localizer.LocalizeKey:
		return NewLocalizeKeyParam(t.LocalizeKey())
	default:
		return NewStringParam(fmt.Sprint(p))
	}
}

// AnyToErrorParams ErrorParamに変換
func AnyToErrorParams(any []*any.Any) ([]ErrorParam, error) {
	result := make([]ErrorParam, len(any))
	for i, a := range any {
		da := &ptypes.DynamicAny{}
		err := ptypes.UnmarshalAny(a, da)
		if err != nil {
			return nil, err
		}
		msgType := reflect.TypeOf(da.Message)

		paramType, ok := messageToParamMap[msgType]
		if !ok {
			return nil, errors.Errorf("no match message type: %v", msgType)
		}

		it := reflect.New(paramType.Elem()).Interface()

		param := it.(ErrorParam)

		if err := param.SetValue(da.Message); err != nil {
			return nil, err
		}

		result[i] = param
	}

	return result, nil
}

// ToAny Anyメッセージに変換
func ToAny(params []ErrorParam) ([]*any.Any, error) {
	if len(params) == 0 {
		return nil, nil
	}

	anyList := make([]*any.Any, len(params))
	for i, p := range params {
		m := p.ToProtoMessage()
		a, err := ptypes.MarshalAny(m)
		if err != nil {
			return nil, err
		}
		anyList[i] = a
	}

	return anyList, nil
}
