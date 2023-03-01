package logger

import (
	"go.uber.org/zap"
)

var Logger appLogger

type appLogger struct {
	Instance *zap.SugaredLogger
}

func InitAppLogger() {
	logger, _ := zap.NewProduction()
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
