package sszap

import (
	"go.uber.org/zap/zapcore"
)

func newDeviceEventEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "date_time",
		LevelKey:       "level",
		NameKey:        "context",
		CallerKey:      zapcore.OmitKey,
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "event_message",
		StacktraceKey:  zapcore.OmitKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    DigitLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func DigitLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var intLevel int
	switch l {
	case zapcore.FatalLevel:
		intLevel = 0
	case zapcore.PanicLevel:
		intLevel = 1
	case zapcore.DPanicLevel:
		intLevel = 2
	case zapcore.ErrorLevel:
		intLevel = 3
	case zapcore.WarnLevel:
		intLevel = 4
	case zapcore.InfoLevel:
		intLevel = 6
	case zapcore.DebugLevel:
		intLevel = 7
	}
	enc.AppendInt(intLevel)
}
