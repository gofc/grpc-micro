package validations

import (
	"github.com/gofc/grpc-micro/code"
	"google.golang.org/grpc/codes"
)

// Result チェック結果
type Result struct {
	ok      bool
	details Details
}

// NewResult チェック処理を開始する
func NewResult() *Result {
	return &Result{
		ok:      true,
		details: NewDetails(),
	}
}

// AddDetail エラー詳細を追加する
func (v *Result) AddDetail(errorCode code.Code, params ...interface{}) *Result {
	d := NewDetail(errorCode, params...)
	v.details = v.details.AddDetail(d)

	if !v.details.IsEmpty() {
		v.ok = false
	}
	return v
}

// Merge チェック結果をマージする
func (v *Result) Merge(other *Result) *Result {
	if other == nil {
		return v
	}
	if !other.ok {
		v.ok = false
	}
	v.details = v.details.Merge(other.details)
	return v
}

// Error エラー詳細からエラーを生成して返却する
func (v *Result) Error() error {
	if v.ok || v.details.IsEmpty() {
		return nil
	}
	return v.details.GenerateError(codes.InvalidArgument)
}

// IsOK チェックOKかどうか
// @return true: OK, false: NG
func (v *Result) IsOK() bool {
	return v.ok
}
