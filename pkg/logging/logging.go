package logging

import (
	"context"
	"go.uber.org/zap"
)

const LOGGER_KEY = 1

var logger *zap.Logger

func initLogger() {
	logger, _ = zap.NewDevelopment()
}

func NewContext(ctx context.Context, fields ...zap.Field) context.Context {
	if logger == nil {
		initLogger()
	}
	return context.WithValue(ctx, LOGGER_KEY, WithContext(ctx).With(fields...))
}

func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return logger
	}
	if ctxLogger, ok := ctx.Value(LOGGER_KEY).(*zap.Logger); ok {
		return ctxLogger
	}
	return logger
}
