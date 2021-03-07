package logging

import (
	"context"
	"go.uber.org/zap"
)

const (
	LoggerKey = 1
)

var logger *zap.Logger

func initLogger() {
	cfg := zap.NewDevelopmentConfig()
	cfg.DisableStacktrace = true
	logger, _ = cfg.Build()
}

func NewContext(ctx context.Context, fields ...zap.Field) context.Context {
	if logger == nil {
		initLogger()
	}
	return context.WithValue(ctx, LoggerKey, WithContext(ctx).With(fields...))
}

func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return logger
	}
	if ctxLogger, ok := ctx.Value(LoggerKey).(*zap.Logger); ok {
		return ctxLogger
	}
	return logger
}
