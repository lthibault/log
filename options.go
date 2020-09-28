package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Option for Logger
type Option func(*logrus.Logger)

// WithLevel sets the log level
func WithLevel(l Level) Option {
	return func(log *logrus.Logger) {
		log.Level = l.toLogrus()
	}
}

// WithFormatter sets the formatter
func WithFormatter(f logrus.Formatter) Option {
	if f == nil {
		f = &logrus.TextFormatter{}
	}

	return func(log *logrus.Logger) {
		log.Formatter = f
	}
}

// WithLevelHooks sets the level hooks
func WithLevelHooks(hs LevelHooks) Option {
	if hs == nil {
		hs = make(LevelHooks)
	}

	return func(log *logrus.Logger) {
		log.Hooks = hs
	}
}

// WithWriter writer
func WithWriter(w io.Writer) Option {
	if w == nil {
		w = os.Stderr
	}

	return func(log *logrus.Logger) {
		log.Out = w
	}
}

func withDefaults(opt []Option) []Option {
	return append([]Option{
		WithLevel(InfoLevel),
	}, opt...)
}
