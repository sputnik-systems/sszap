package sszap

import (
	"runtime"
	"time"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type stdoutEncoder struct {
	*zapcore.EncoderConfig
	internal zapcore.Encoder
}

func (enc stdoutEncoder) AddArray(key string, arr zapcore.ArrayMarshaler) error {
	return enc.internal.AddArray(key, arr)
}

func (enc stdoutEncoder) AddObject(key string, obj zapcore.ObjectMarshaler) error {
	return enc.internal.AddObject(key, obj)
}

func (enc stdoutEncoder) AddBinary(key string, val []byte) {
	enc.internal.AddBinary(key, val)
}

func (enc stdoutEncoder) AddByteString(key string, val []byte) {
	enc.internal.AddByteString(key, val)
}

func (enc stdoutEncoder) AddBool(key string, val bool) {
	enc.internal.AddBool(key, val)
}

func (enc stdoutEncoder) AddComplex128(key string, val complex128) {
	enc.internal.AddComplex128(key, val)
}

func (enc stdoutEncoder) AddComplex64(key string, val complex64) {
	enc.internal.AddComplex64(key, val)
}

func (enc stdoutEncoder) AddDuration(key string, val time.Duration) {
	enc.internal.AddDuration(key, val)
}

func (enc stdoutEncoder) AddFloat64(key string, val float64) {
	enc.internal.AddFloat64(key, val)
}

func (enc stdoutEncoder) AddFloat32(key string, val float32) {
	enc.internal.AddFloat32(key, val)
}

func (enc stdoutEncoder) AddInt(key string, val int) {
	enc.internal.AddInt(key, val)
}

func (enc stdoutEncoder) AddInt64(key string, val int64) {
	enc.internal.AddInt64(key, val)
}

func (enc stdoutEncoder) AddInt32(key string, val int32) {
	enc.internal.AddInt32(key, val)
}

func (enc stdoutEncoder) AddInt16(key string, val int16) {
	enc.internal.AddInt16(key, val)
}

func (enc stdoutEncoder) AddInt8(key string, val int8) {
	enc.internal.AddInt8(key, val)
}

func (enc stdoutEncoder) AddString(key, val string) {
	enc.internal.AddString(key, val)
}

func (enc stdoutEncoder) AddTime(key string, val time.Time) {
	enc.internal.AddTime(key, val)
}

func (enc stdoutEncoder) AddUint(key string, val uint) {
	enc.internal.AddUint(key, val)
}

func (enc stdoutEncoder) AddUint64(key string, val uint64) {
	enc.internal.AddUint64(key, val)
}

func (enc stdoutEncoder) AddUint32(key string, val uint32) {
	enc.internal.AddUint32(key, val)
}

func (enc stdoutEncoder) AddUint16(key string, val uint16) {
	enc.internal.AddUint16(key, val)
}

func (enc stdoutEncoder) AddUint8(key string, val uint8) {
	enc.internal.AddUint8(key, val)
}

func (enc stdoutEncoder) AddUintptr(key string, val uintptr) {
	enc.internal.AddUintptr(key, val)
}

func (enc stdoutEncoder) AddReflected(key string, val interface{}) error {
	return enc.internal.AddReflected(key, val)
}

func (enc stdoutEncoder) OpenNamespace(key string) {
	enc.internal.OpenNamespace(key)
}

func (enc stdoutEncoder) Clone() zapcore.Encoder {
	return stdoutEncoder{
		EncoderConfig: enc.EncoderConfig,
		internal:      enc.internal.Clone(),
	}
}

func (enc stdoutEncoder) EncodeEntry(ent zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	if ent.Level > zapcore.WarnLevel {
		rl := sourceLocationFromEntry(ent)
		if rl != nil {
			fields = append(fields, SourceLocationField(rl))
		}
	}

	return enc.internal.EncodeEntry(ent, fields)
}

func sourceLocationFromEntry(ent zapcore.Entry) *SourceLocation {
	caller := ent.Caller
	if !caller.Defined {
		return nil
	}
	loc := &SourceLocation{
		FilePath:   caller.File,
		LineNumber: caller.Line,
	}
	if fn := runtime.FuncForPC(caller.PC); fn != nil {
		loc.FunctionName = fn.Name()
	}

	return loc
}

func newStdoutEncoder(cfg zapcore.EncoderConfig) zapcore.Encoder {
	return stdoutEncoder{
		EncoderConfig: &cfg,
		internal:      zapcore.NewJSONEncoder(cfg),
	}
}

func newStdoutEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "severity",
		NameKey:        zapcore.OmitKey,
		CallerKey:      zapcore.OmitKey,
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "text_payload",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    SeverityLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

func SeverityLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(logLevelSeverity[l])
}
