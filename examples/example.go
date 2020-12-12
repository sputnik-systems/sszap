package examples

import (
	"context"

	"github.com/sputnik-systems/sszap"
	"go.uber.org/zap"
)

func simpleLogger() { //nolint
	logLevel := "debug"

	logger := sszap.NewLogger(
		sszap.NewPreparedStdoutCore(logLevel),
	)
	sszap.SetDefaultLogger(logger.Named("test"))

	ctx := context.Background()
	ctxLogger := sszap.FromContext(ctx)

	ctxLogger.With(
		zap.String("test_field", "test_value"),
	).Info("New info message")
}
