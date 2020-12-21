package logger

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger
var noStackLogger *zap.Logger

//Config ログ設定
type Config struct {
	Level  string `yaml:"level" default:"info"`
	Format string `yaml:"format" default:"json"`
}

type noStackLevelEnabler struct{}

//Enabled Stack出力無効化
func (le *noStackLevelEnabler) Enabled(zapcore.Level) bool {
	return false
}

//InitLogger ログ初期化する
func InitLogger(config *Config) {
	atomicLevel := zap.NewAtomicLevel()
	if atomicLevel.UnmarshalText([]byte(config.Level)) != nil {
		atomicLevel.SetLevel(zap.DebugLevel)
	}
	level := atomicLevel.Level()

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel && lvl >= level
	})

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	jsonEncoder := zapcore.NewJSONEncoder(encoderCfg)
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	var cores = make([]zapcore.Core, 0)
	switch config.Format {
	case "text":
		cores = append(cores,
			zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
			zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority))
	case "json":
		cores = append(cores,
			zapcore.NewCore(jsonEncoder, consoleErrors, highPriority),
			zapcore.NewCore(jsonEncoder, consoleDebugging, lowPriority))
	default:
		panic(fmt.Sprintf("invalid format value:%v", config.Format))
	}

	core := zapcore.NewTee(cores...)

	logger = zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel))
	noStackLogger = logger.WithOptions(zap.AddStacktrace(&noStackLevelEnabler{}))
}

//Close ログ出力をクローズする
func Close(ctx context.Context) {
	err := logger.Sync()
	if err != nil {
		fmt.Println("logger.Sync error", err)
	}
}

//GetLogger logger対象を取得する
func GetLogger() *zap.Logger {
	return logger
}
