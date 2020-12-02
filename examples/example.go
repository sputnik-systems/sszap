package examples

import (
	"context"

	"github.com/sputnik-systems/sszap"
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
		sszap.DeviceIDField("test_id"),
	).Info("New info message")
}
