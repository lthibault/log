package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

type key uint16

const (
	// FatalLevel logs and then calls `os.Exit(1)`.
	FatalLevel Level = logrus.FatalLevel
	// ErrorLevel is used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel Level = logrus.ErrorLevel
	// WarnLevel is for non-critical entries that deserve eyes.
	WarnLevel Level = logrus.WarnLevel
	// InfoLevel provides general operational entries about what's going on inside the
	// application.
	InfoLevel Level = logrus.InfoLevel
	// DebugLevel is used to report application state for debugging perposes.
	DebugLevel Level = logrus.DebugLevel
	// TraceLevel is used to trace the execution steps of an application for debugging
	// or optimization purposes.
	TraceLevel Level = logrus.TraceLevel
)

type (
	// LevelHooks .
	LevelHooks = logrus.LevelHooks
	// Level .
	Level = logrus.Level
)

type spec struct {
	Hooks     LevelHooks
	Formatter logrus.Formatter
	Level     Level
	io.Writer
}

func (s spec) mkLogger() Logger {
	l := logrus.New()
	l.SetLevel(s.Level)

	if s.Hooks != nil {
		l.Hooks = s.Hooks
	}

	if s.Formatter != nil {
		l.Formatter = s.Formatter
	}

	if s.Writer != nil {
		l.Out = s.Writer
	}

	return (*entry)(logrus.NewEntry(l))
}

// Option for Logger
type Option func(*spec) Option

// WithLevel sets the log level
func WithLevel(l Level) Option {
	return func(c *spec) (prev Option) {
		prev = WithLevel(c.Level)
		c.Level = l
		return
	}
}

// WithFormatter sets the formatter
func WithFormatter(f logrus.Formatter) Option {
	return func(c *spec) (prev Option) {
		prev = WithFormatter(c.Formatter)
		c.Formatter = f
		return
	}
}

// WithLevelHooks sets the level hooks
func WithLevelHooks(h LevelHooks) Option {
	return func(c *spec) (prev Option) {
		prev = WithLevelHooks(c.Hooks)
		c.Hooks = h
		return
	}
}

// WithWriter writer
func WithWriter(w io.Writer) Option {
	return func(c *spec) (prev Option) {
		prev = WithWriter(c.Writer)
		c.Writer = w
		return
	}
}

func withDefaults(opt []Option) []Option {
	return append([]Option{
		WithLevel(InfoLevel),
	}, opt...)
}
