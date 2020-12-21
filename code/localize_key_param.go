package code

import (
	"fmt"
	"reflect"

	"github.com/gofc/grpc-micro/localizer"
	pbbase "github.com/gofc/grpc-micro/proto"

	"github.com/golang/protobuf/proto"
)

// LocalizeKeyParam ローカライズ用のを格納したエラーパラメータ
type LocalizeKeyParam struct {
	value string
}

// NewLocalizeKeyParam 新規作成
func NewLocalizeKeyParam(value string) *LocalizeKeyParam {
	return &LocalizeKeyParam{
		value: value,
	}
}

// Value valueを返す
func (p *LocalizeKeyParam) Value() string {
	return p.value
}

// Localize ローカライズ
func (p *LocalizeKeyParam) Localize(localizer localizer.Localizer) string {
	s, err := localizer.Localize(p.value, nil)
	if err != nil {
		return p.value
	}
	return s
}

// ToProtoMessage protoメッセージを作成
func (p *LocalizeKeyParam) ToProtoMessage() proto.Message {
	return &pbbase.MessageIDParam{
		Value: p.value,
	}
}

// MessageType protoメッセージの型
func (p *LocalizeKeyParam) MessageType() reflect.Type {
	return reflect.TypeOf(&pbbase.MessageIDParam{})
}

// SetValue protoメッセージから値を詰める
func (p *LocalizeKeyParam) SetValue(msg proto.Message) error {
	m := msg.(*pbbase.MessageIDParam)
	p.value = m.Value
	return nil
}

func (p *LocalizeKeyParam) String() string {
	return fmt.Sprintf("LocalizeKey{%s}", p.value)
}
