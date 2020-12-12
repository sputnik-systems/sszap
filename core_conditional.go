package sszap

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// conditionalCore writes log entry only if activator field passed.
type conditionalCore struct {
	zapcore.LevelEnabler
	enc            zapcore.Encoder
	out            zapcore.WriteSyncer
	activated      bool
	activatorField string
}

func newConditionalCore(enc zapcore.Encoder, ws zapcore.WriteSyncer,
	enab zapcore.LevelEnabler, activator string) zapcore.Core {
	return &conditionalCore{
		LevelEnabler:   enab,
		enc:            enc,
		out:            ws,
		activatorField: activator,
	}
}

func (c *conditionalCore) With(fields []zapcore.Field) zapcore.Core {
	clone := c.clone()
	for _, f := range fields {
		if f.Key == c.activatorField {
			clone.activated = true

			continue
		}
		f.AddTo(clone.enc)
	}

	return clone
}

func (c *conditionalCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if c.Enabled(ent.Level) && c.activated {
		return ce.AddCore(ent, c)
	}

	return ce
}

func (c *conditionalCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
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
		c.Sync() // nolint
	}

	return nil
}

func (c *conditionalCore) Sync() error {
	return c.out.Sync()
}

func (c *conditionalCore) clone() *conditionalCore {
	return &conditionalCore{
		LevelEnabler: c.LevelEnabler,
		enc:          c.enc.Clone(),
		out:          c.out,
		activated:    c.activated,
	}
}
