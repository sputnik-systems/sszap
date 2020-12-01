package examples

import (
	"context"
	"fmt"

	"go.uber.org/zap/zapcore"

	"github.com/sputnik-systems/sszap"
)

type eWriteSyncer struct {
}

var _ zapcore.WriteSyncer = (*eWriteSyncer)(nil)

func (ws *eWriteSyncer) Write(b []byte) (int, error) {
	fmt.Println(string(b))
	return len(b), nil
}

func (ws *eWriteSyncer) Sync() error {
	return nil
}

func deviceLogger() { //nolint
	logLevel := "debug"

	logger := sszap.NewLogger(
		sszap.NewPreparedDeviceCore(logLevel, &eWriteSyncer{}),
	)
	sszap.SetDefaultLogger(logger)

	ctx := context.Background()

	ctxLogger := sszap.FromContext(ctx)

	ctxLogger.With(
		sszap.DeviceIDField("test_id"),
	).Info("New info message")

}
