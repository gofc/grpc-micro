package code

import (
	"fmt"
	"reflect"
	"time"

	"github.com/gofc/grpc-micro/localizer"
	pbbase "github.com/gofc/grpc-micro/proto"

	"github.com/golang/protobuf/proto"
)

// TimeParam ローカライズ用のを格納したエラーパラメータ
type TimeParam struct {
	value  time.Time
	format string
}

// NewTimeParam 新規作成
func NewTimeParam(value time.Time, format string) *TimeParam {
	return &TimeParam{
		value:  value,
		format: format,
	}
}

// Value valueを返す
func (p *TimeParam) Value() string {
	return p.value.Format(p.format)
}

// Localize ローカライズ
func (p *TimeParam) Localize(localizer localizer.Localizer) string {
	return p.value.In(localizer.Location()).Format(p.format)
}

// ToProtoMessage protoメッセージを作成
func (p *TimeParam) ToProtoMessage() proto.Message {
	return &pbbase.TimeParam{
		Value:  p.value.Unix(),
		Format: p.format,
	}
}

// MessageType protoメッセージの型
func (p *TimeParam) MessageType() reflect.Type {
	return reflect.TypeOf(&pbbase.TimeParam{})
}

// SetValue protoメッセージから値を詰める
func (p *TimeParam) SetValue(msg proto.Message) error {
	m := msg.(*pbbase.TimeParam)
	p.value = time.Unix(m.Value, 0)
	p.format = m.Format
	return nil
}

func (p *TimeParam) String() string {
	return fmt.Sprintf("Time{%s,%s}", p.value, p.format)
}
