package logger

import (
	"context"

	"go.uber.org/zap"
)

type loggerCtx struct{}

var nopLogger = zap.NewNop()

func Context(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerCtx{}, logger)
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	fromContext(ctx).Debug(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	fromContext(ctx).Info(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	fromContext(ctx).Warn(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	fromContext(ctx).Error(msg, fields...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	fromContext(ctx).Panic(msg, fields...)
}

func fromContext(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(loggerCtx{}).(*zap.Logger); ok {
		return logger
	}

	return nopLogger
}
