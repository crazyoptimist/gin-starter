package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger appLogger

type appLogger struct {
	Instance *zap.SugaredLogger
}

func InitAppLogger() {

	config := zap.NewProductionConfig()

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	// Because we want to see the actual caller, we skip caller's call stack depth by one
	logger, _ := config.Build(zap.AddCallerSkip(1))

	sugar := logger.Sugar()

	Logger = appLogger{Instance: sugar}
}

func (l *appLogger) Info(data interface{}) {
	l.Instance.Info(data)
}

func (l *appLogger) Warn(data interface{}) {
	l.Instance.Warn(data)
}

func (l *appLogger) Error(data interface{}) {
	l.Instance.Error(data)
}
