package sszap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const sourceKey = "source_location"

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
