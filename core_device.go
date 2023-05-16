package sszap

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// deviceEventCore writes log entry only if event field is setted.
type deviceEventCore struct {
	zapcore.LevelEnabler
	enc            zapcore.Encoder
	out            zapcore.WriteSyncer
	hasDeviceEvent bool
}

func newDeviceEventCore(enc zapcore.Encoder, ws zapcore.WriteSyncer, enab zapcore.LevelEnabler) zapcore.Core {
	return &deviceEventCore{
		LevelEnabler: enab,
		enc:          enc,
		out:          ws,
	}
}

func (c *deviceEventCore) With(fields []zapcore.Field) zapcore.Core {
	clone := c.clone()
	for _, f := range fields {
		if f.Key == eventKey {
			clone.hasDeviceEvent = true
		}
		f.AddTo(clone.enc)
	}

	return clone
}

func (c *deviceEventCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) && c.hasDeviceEvent {
		return ce.AddCore(ent, c)
	}

	return ce
}

func (c *deviceEventCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	buf, err := c.enc.EncodeEntry(ent, fields)
	if err != nil {
		return fmt.Errorf("core encode error: %w", err)
	}
	_, err = c.out.Write(buf.Bytes())
	buf.Free()
	if err != nil {
		return fmt.Errorf("core write error: %w", err)
	}
	if ent.Level > zapcore.ErrorLevel {
		// Since we may be crashing the program, sync the output. Ignore Sync
		// errors, pending a clean solution to issue #370.
		c.Sync() //nolint
	}

	return nil
}

func (c *deviceEventCore) Sync() error {
	return c.out.Sync()
}

func (c *deviceEventCore) clone() *deviceEventCore {
	return &deviceEventCore{
		LevelEnabler:   c.LevelEnabler,
		enc:            c.enc.Clone(),
		out:            c.out,
		hasDeviceEvent: c.hasDeviceEvent,
	}
}
