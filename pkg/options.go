package log

import (
	"io"
	"unsafe"

	"github.com/sirupsen/logrus"
)

type key uint16

const (
	keyLogger key = iota

	// NullLevel disables all logging
	NullLevel Level = logrus.PanicLevel
	// FatalLevel level. Logs and then calls `os.Exit(1)`.
	FatalLevel Level = logrus.FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel Level = logrus.ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel Level = logrus.WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel Level = logrus.InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel Level = logrus.DebugLevel
)

type (
	LevelHooks = logrus.LevelHooks
	Formatter  = logrus.Formatter
	Level      = logrus.Level
)

// Cfg creates a Logger using the logrus library
type Cfg struct {
	Hooks     LevelHooks
	Formatter Formatter
	Level     Level
	io.Writer
}

func (cfg Cfg) mkLogger() Logger {
	if cfg.Level == NullLevel {
		return noop{}
	}

	l := logrus.New()
	l.SetLevel(cfg.Level)

	if cfg.Hooks != nil {
		l.Hooks = cfg.Hooks
	}

	if cfg.Formatter != nil {
		l.Formatter = cfg.Formatter
	}

	if cfg.Writer != nil {
		l.Out = cfg.Writer
	}

	return (*entry)(unsafe.Pointer(logrus.NewEntry(l)))
}

// Option for Logger
type Option func(*Cfg) Option

// OptLevel sets the log level
func OptLevel(l Level) Option {
	return func(c *Cfg) (prev Option) {
		prev = OptLevel(c.Level)
		c.Level = l
		return
	}
}

// OptFormatter sets the formatter
func OptFormatter(f Formatter) Option {
	return func(c *Cfg) (prev Option) {
		prev = OptFormatter(c.Formatter)
		c.Formatter = f
		return
	}
}

// OptLevelHooks sets the level hooks
func OptLevelHooks(h LevelHooks) Option {
	return func(c *Cfg) (prev Option) {
		prev = OptLevelHooks(c.Hooks)
		c.Hooks = h
		return
	}
}

// OptWriter writer
func OptWriter(w io.Writer) Option {
	return func(c *Cfg) (prev Option) {
		prev = OptWriter(c.Writer)
		c.Writer = w
		return
	}
}
