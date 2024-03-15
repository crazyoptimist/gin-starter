package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TODO: Replace zap with Go 1.22 structured logging
var Logger appLogger

type appLogger struct {
	Instance *zap.SugaredLogger
}

func InitAppLogger() (*appLogger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	// Because we want to see the actual caller, we skip caller's call stack depth by one
	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil, err
	}

	sugar := logger.Sugar()

	Logger = appLogger{Instance: sugar}

	return &Logger, err
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
