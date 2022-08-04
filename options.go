package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Option for Logger
type Option func(*logrus.Logger)

// WithLevel sets the log level
func WithLevel(lvl Level) Option {
	return func(log *logrus.Logger) {
		switch lvl {
		case FatalLevel:
			log.Level = logrus.FatalLevel
		case ErrorLevel:
			log.Level = logrus.ErrorLevel
		case WarnLevel:
			log.Level = logrus.WarnLevel
		case InfoLevel:
			log.Level = logrus.InfoLevel
		case DebugLevel:
			log.Level = logrus.DebugLevel
		case TraceLevel:
			log.Level = logrus.TraceLevel
		default:
			panic(lvl)
		}
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
