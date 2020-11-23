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

	sszap.InitLogger(
		sszap.NewPreparedDeviceCore(logLevel, &eWriteSyncer{}),
	)

	ctx := context.Background()

	logger := sszap.FromContext(ctx)

	logger.With(
		sszap.DeviceIDField("test_id"),
	).Info("New info message")

}
