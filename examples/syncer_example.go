package examples

import (
	"context"
	"fmt"

	"github.com/sputnik-systems/sszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
		sszap.NewConditionalCore(logLevel, activatorField, &eWriteSyncer{}),
	)
	sszap.SetDefaultLogger(logger)

	ctx := context.Background()

	ctxLogger := sszap.FromContext(ctx)

	ctxLogger.With(
		GlobalLog(),
		zap.String("additional", "some very important info"),
	).Info("New info message")
}

const activatorField = "is_global"

func GlobalLog() zapcore.Field {
	return zap.Bool(activatorField, true)
}
