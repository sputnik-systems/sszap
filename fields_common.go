package sszap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	deviceIDKey  = "device_id"
	eventKey     = "event"
	eventCodeKey = "event_code"
	dataKey      = "data"
	sourceKey    = "source_location"
)

func DeviceIDField(e string) zapcore.Field {
	return zap.String(deviceIDKey, e)
}

func EventField(e string) zapcore.Field {
	return zap.String(eventKey, e)
}

func EventCodeField(c int) zapcore.Field {
	return zap.Int(eventCodeKey, c)
}

type SourceLocation struct {
	FilePath     string
	LineNumber   int
	FunctionName string
}

func (sl *SourceLocation) Clone() *SourceLocation {
	return &SourceLocation{
		FilePath:     sl.FilePath,
		LineNumber:   sl.LineNumber,
		FunctionName: sl.FunctionName,
	}
}

func (sl SourceLocation) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("file", sl.FilePath)
	enc.AddInt("line", sl.LineNumber)
	enc.AddString("function", sl.FunctionName)

	return nil
}

func SourceLocationField(sl *SourceLocation) zapcore.Field {
	return zap.Object(sourceKey, sl)
}

type ConnectionData struct {
	IsOnline bool
	Sid      string
}

func (sc *ConnectionData) Clone() *ConnectionData {
	return &ConnectionData{
		IsOnline: sc.IsOnline,
		Sid:      sc.Sid,
	}
}

func (sc *ConnectionData) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddBool("is_online", sc.IsOnline)
	enc.AddString("sid", sc.Sid)

	return nil
}

func ConnectionDataField(sc *ConnectionData) zapcore.Field {
	return zap.Object(dataKey, sc)
}
