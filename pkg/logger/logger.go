package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

//Debugf レベルログを出力
func Debugf(format string, args ...interface{}) {
	if logger.Core().Enabled(zap.DebugLevel) {
		if len(args) > 0 {
			format = fmt.Sprintf(format, args...)
		}
		logger.Debug(format)
	}
}

//CDebugf レベルログを出力
func CDebugf(ctx context.Context, format string, args ...interface{}) {
	if logger.Core().Enabled(zap.DebugLevel) {
		if len(args) > 0 {
			format = fmt.Sprintf(format, args...)
		}
		logger.Debug(format, GetContextFields(ctx))
	}
}

//CInfo レベルログを出力
func CInfo(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, GetContextFields(ctx))
	logger.Info(msg, fields...)
}

//Error レベルログを出力
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

//CError レベルログを出力
func CError(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, GetContextFields(ctx))
	logger.Error(msg, fields...)
}

//CErrorNoStack Stack出力なしで Error レベルログを出力
func CErrorNoStack(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, GetContextFields(ctx))
	noStackLogger.Error(msg, fields...)
}

//GetContextFields contextの内容抽出
func GetContextFields(ctx context.Context) zap.Field {
	return zap.Skip()
	//if ctx == nil {
	//	return zap.Skip()
	//}
	//rid := contexts.GetRequestID(ctx)
	//if rid == "" {
	//	return zap.Skip()
	//}
	//return zap.String("requestID", rid)
}
