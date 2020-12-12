package sszap

import (
	"context"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerKey struct{}

var fallbackLogger *zap.SugaredLogger

func init() {
	logger, _ := zap.NewProduction()
	fallbackLogger = logger.Sugar()
}

func SetDefaultLogger(zl *zap.SugaredLogger) {
	fallbackLogger = zl
}

func NewLogger(cores ...zapcore.Core) *zap.SugaredLogger {
	return zap.New(zapcore.NewTee(cores...), zap.AddCaller()).Sugar()
}

func NewPreparedStdoutCore(level string) zapcore.Core {
	return zapcore.NewCore(
		newStdoutEncoder(newStdoutEncoderConfig()),
		zapcore.Lock(os.Stdout),
		levelEnabler(level),
	)
}

func levelEnabler(l string) zap.LevelEnablerFunc {
	return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= parseLogLevel(l)
	})
}

func parseLogLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case zapcore.ErrorLevel.String():
		return zapcore.ErrorLevel
	case zapcore.WarnLevel.String():
		return zapcore.WarnLevel
	case zapcore.DebugLevel.String():
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}

func NewConditionalCore(level, activator string, ws zapcore.WriteSyncer) zapcore.Core {
	return newConditionalCore(
		zapcore.NewJSONEncoder(newStdoutEncoderConfig()),
		ws,
		levelEnabler(level),
		activator,
	)
}

func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	if logger, ok := ctx.Value(loggerKey{}).(*zap.SugaredLogger); ok {
		return logger
	}

	return fallbackLogger
}
