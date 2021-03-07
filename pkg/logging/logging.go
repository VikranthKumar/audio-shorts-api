package logging

import (
	"context"
	"go.uber.org/zap"
)

const (
	LoggerKey = 1
)

var logger *zap.Logger

// initLogger creates the context
func initLogger() {
	cfg := zap.NewDevelopmentConfig()
	cfg.DisableStacktrace = true
	logger, _ = cfg.Build()
}

// NewContext provides a context with logger
func NewContext(ctx context.Context, fields ...zap.Field) context.Context {
	if logger == nil {
		initLogger()
	}
	return context.WithValue(ctx, LoggerKey, WithContext(ctx).With(fields...))
}

// WithContext provides a logger from the context
func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return logger
	}
	if ctxLogger, ok := ctx.Value(LoggerKey).(*zap.Logger); ok {
		return ctxLogger
	}
	return logger
}
