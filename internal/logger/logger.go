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

func (l *appLogger) Debug(args ...interface{}) {
	l.Instance.Debug(args)
}

func (l *appLogger) Info(args ...interface{}) {
	l.Instance.Info(args)
}

func (l *appLogger) Warn(args ...interface{}) {
	l.Instance.Warn(args)
}

func (l *appLogger) Error(args ...interface{}) {
	l.Instance.Error(args)
}

func (l *appLogger) Fatal(args ...interface{}) {
	l.Instance.Fatal(args)
}
