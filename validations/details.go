package validations

import (
	"fmt"
	"github.com/gofc/grpc-micro/code"
	"github.com/gofc/grpc-micro/localizer"
	pbbase "github.com/gofc/grpc-micro/proto"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"strconv"
)

// Detail エラー詳細(単体)
type Detail struct {
	ErrorCode code.Code
	Params    []code.ErrorParam
}

// NewDetail エラー詳細(単体)を生成
func NewDetail(messageID code.Code, ps ...interface{}) *Detail {
	detail := &Detail{
		ErrorCode: messageID,
	}
	if ps != nil {
		detail.Params = code.ToErrorParams(ps)
	}
	return detail
}

// String エラー詳細(単数)を文字列にして返却する
func (detail *Detail) String() string {
	return fmt.Sprintf("error_detail:[%+v, params:%+v]", detail.ErrorCode, detail.Params)
}

// Localize localize
func (detail *Detail) Localize(l localizer.Localizer) string {
	data := make(map[string]interface{}, len(detail.Params))
	for i, p := range detail.Params {
		key := "Param" + strconv.Itoa(i)
		data[key] = p.Localize(l)
	}

	s, err := l.Localize(detail.ErrorCode.LocalizeKey(), data)
	if err != nil {
		return detail.ErrorCode.LocalizeKey()
	}

	return s
}

// ToProtoMessage メッセージに変換
func (detail *Detail) ToProtoMessage() (*pbbase.ValidationError, error) {
	any, err := code.ToAny(detail.Params)
	if err != nil {
		return nil, err
	}

	return &pbbase.ValidationError{
		Code:   uint64(detail.ErrorCode),
		Params: any,
	}, nil
}

// Details エラー詳細(複数)
type Details []*Detail

// NewDetails エラー詳細を生成
func NewDetails() Details {
	return *new(Details)
}

// AddDetail エラー詳細を追加
func (details Details) AddDetail(detail *Detail) Details {
	if detail.ErrorCode == 0 {
		return details
	}
	return append(details, detail)
}

// Get Indexを指定してエラー詳細を取得する
func (details Details) Get(index int) *Detail {
	if index < 0 || len(details) <= index {
		return nil
	}
	return details[index]
}

// Merge エラー詳細をマージする
func (details Details) Merge(other Details) Details {
	return append(details, other...)
}

// IsEmpty 空かどうか
// @return true: empty, false: not empty
func (details Details) IsEmpty() bool {
	return len(details) == 0
}

// GenerateError エラー詳細(Details)を含めたエラーを生成する
func (details Details) GenerateError(code codes.Code) error {
	if code == codes.OK && details.IsEmpty() {
		return nil
	}

	st := status.New(code, details.String())
	ds := make([]proto.Message, 0)

	for _, d := range details {
		if msgDetail, err := d.ToProtoMessage(); err == nil {
			ds = append(ds, msgDetail)
		}
	}
	st, err := st.WithDetails(ds...)
	if err != nil {
		return err
	}
	return st.Err()
}

// String エラー詳細(複数)を文字列にして返却する
func (details Details) String() string {
	if !details.IsEmpty() {
		msgs := make([]string, 0)
		for _, d := range details {
			msgs = append(msgs, d.String())
		}
		return fmt.Sprintf("error_details:%+v", msgs)
	}
	return "error_details:Nothing"
}
