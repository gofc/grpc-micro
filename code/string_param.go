package code

import (
	"reflect"

	"github.com/gofc/grpc-micro/localizer"
	pbbase "github.com/gofc/grpc-micro/proto"

	"github.com/golang/protobuf/proto"
)

// StringParam 文字列
type StringParam struct {
	value string
}

// NewStringParam 新規作成
func NewStringParam(value string) *StringParam {
	return &StringParam{
		value: value,
	}
}

// Value valueを返す
func (p *StringParam) Value() string {
	return p.value
}

// Localize ローカライズ
func (p *StringParam) Localize(localizer localizer.Localizer) string {
	return p.value
}

// ToProtoMessage protoメッセージを作成
func (p *StringParam) ToProtoMessage() proto.Message {
	return &pbbase.StringParam{
		Value: p.value,
	}
}

// MessageType protoメッセージの型
func (p *StringParam) MessageType() reflect.Type {
	return reflect.TypeOf(&pbbase.StringParam{})
}

// SetValue protoメッセージから値を詰める
func (p *StringParam) SetValue(msg proto.Message) error {
	m := msg.(*pbbase.StringParam)
	p.value = m.Value
	return nil
}

func (p *StringParam) String() string {
	return p.value
}
