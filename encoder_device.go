package sszap

import (
	"time"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type deviceEventEncoder struct {
	*zapcore.EncoderConfig
	internal zapcore.Encoder
}

func (enc deviceEventEncoder) AddArray(key string, arr zapcore.ArrayMarshaler) error {
	return enc.internal.AddArray(key, arr)
}

func (enc deviceEventEncoder) AddObject(key string, obj zapcore.ObjectMarshaler) error {
	return enc.internal.AddObject(key, obj)
}

func (enc deviceEventEncoder) AddBinary(key string, val []byte) {
	enc.internal.AddBinary(key, val)
}

func (enc deviceEventEncoder) AddByteString(key string, val []byte) {
	enc.internal.AddByteString(key, val)
}

func (enc deviceEventEncoder) AddBool(key string, val bool) {
	enc.internal.AddBool(key, val)
}

func (enc deviceEventEncoder) AddComplex128(key string, val complex128) {
	enc.internal.AddComplex128(key, val)
}

func (enc deviceEventEncoder) AddComplex64(key string, val complex64) {
	enc.internal.AddComplex64(key, val)
}

func (enc deviceEventEncoder) AddDuration(key string, val time.Duration) {
	enc.internal.AddDuration(key, val)
}

func (enc deviceEventEncoder) AddFloat64(key string, val float64) {
	enc.internal.AddFloat64(key, val)
}

func (enc deviceEventEncoder) AddFloat32(key string, val float32) {
	enc.internal.AddFloat32(key, val)
}

func (enc deviceEventEncoder) AddInt(key string, val int) {
	enc.internal.AddInt(key, val)
}

func (enc deviceEventEncoder) AddInt64(key string, val int64) {
	enc.internal.AddInt64(key, val)
}

func (enc deviceEventEncoder) AddInt32(key string, val int32) {
	enc.internal.AddInt32(key, val)
}

func (enc deviceEventEncoder) AddInt16(key string, val int16) {
	enc.internal.AddInt16(key, val)
}

func (enc deviceEventEncoder) AddInt8(key string, val int8) {
	enc.internal.AddInt8(key, val)
}

func (enc deviceEventEncoder) AddString(key, val string) {
	enc.internal.AddString(key, val)
}

func (enc deviceEventEncoder) AddTime(key string, val time.Time) {
	enc.internal.AddTime(key, val)
}

func (enc deviceEventEncoder) AddUint(key string, val uint) {
	enc.internal.AddUint(key, val)
}

func (enc deviceEventEncoder) AddUint64(key string, val uint64) {
	enc.internal.AddUint64(key, val)
}

func (enc deviceEventEncoder) AddUint32(key string, val uint32) {
	enc.internal.AddUint32(key, val)
}

func (enc deviceEventEncoder) AddUint16(key string, val uint16) {
	enc.internal.AddUint16(key, val)
}

func (enc deviceEventEncoder) AddUint8(key string, val uint8) {
	enc.internal.AddUint8(key, val)
}

func (enc deviceEventEncoder) AddUintptr(key string, val uintptr) {
	enc.internal.AddUintptr(key, val)
}

func (enc deviceEventEncoder) AddReflected(key string, val interface{}) error {
	return enc.internal.AddReflected(key, val)
}

func (enc deviceEventEncoder) OpenNamespace(key string) {
	enc.internal.OpenNamespace(key)
}

func (enc deviceEventEncoder) Clone() zapcore.Encoder {
	return deviceEventEncoder{
		EncoderConfig: enc.EncoderConfig,
		internal:      enc.internal.Clone(),
	}
}

func (enc deviceEventEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	fields = append(fields, ContextField())

	return enc.internal.EncodeEntry(ent, fields)
}

func newDeviceEventEncoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	return deviceEventEncoder{
		EncoderConfig: &cfg,
		internal:      zapcore.NewJSONEncoder(cfg),
	}
}

func newDeviceEventEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "date_time",
		LevelKey:       "level",
		NameKey:        zapcore.OmitKey,
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
