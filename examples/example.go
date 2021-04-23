package examples

import (
	"context"

	"github.com/sputnik-systems/sszap"
)

func simpleLogger() { //nolint
	logLevel := "debug"

	sszap.InitLogger(
		sszap.NewPreparedStdoutCore(logLevel),
	)

	ctx := context.Background()

	logger := sszap.FromContext(ctx)

	logger.With(
		sszap.DeviceIDField("test_id"),
		sszap.NodeIDField("test_node_id"),
	).Info("New info message")

}
