package logger

import (
	"context"

	"go.uber.org/zap"
)

type ctxMarker struct{}

var (
	traceKey = &ctxMarker{}
)

// SetTraceFields trace 情報を context に追加する
func SetTraceFields(ctx context.Context, traceID, spanID uint64) context.Context {
	return context.WithValue(ctx, traceKey, []zap.Field{
		zap.Uint64("dd.trace_id", traceID),
		zap.Uint64("dd.span_id", spanID),
	})
}

// GetTraceFields context の trace 情報を zap.Field 配列へ変換し, 戻り値として返す
func GetTraceFields(ctx context.Context) []zap.Field {
	t, ok := ctx.Value(traceKey).([]zap.Field)
	if !ok {
		return nil
	}
	return t
}
