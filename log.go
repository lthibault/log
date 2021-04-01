// Package log contains an improved API and utilities for github.com/sirupsen/logrus
package log

//go:generate mockgen -source=log.go -destination=test/logtest.go -package=logtest

import (
	"github.com/sirupsen/logrus"
)

const (
	// FatalLevel logs and then calls `os.Exit(1)`.
	FatalLevel Level = iota

	// ErrorLevel is used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel

	// WarnLevel is for non-critical entries that deserve eyes.
	WarnLevel

	// InfoLevel provides general operational entries about what's going on inside the
	// application.
	InfoLevel

	// DebugLevel is used to report application state for debugging perposes.
	DebugLevel

	// TraceLevel is used to trace the execution steps of an application for debugging
	// or optimization purposes.
	TraceLevel
)

// Level .
type Level uint8

func (lvl Level) toLogrus() logrus.Level {
	switch lvl {
	case FatalLevel:
		return logrus.FatalLevel
	case ErrorLevel:
		return logrus.ErrorLevel
	case WarnLevel:
		return logrus.WarnLevel
	case InfoLevel:
		return logrus.InfoLevel
	case DebugLevel:
		return logrus.DebugLevel
	case TraceLevel:
		return logrus.TraceLevel
	default:
		panic(lvl)
	}
}

// LevelHooks .
type LevelHooks = logrus.LevelHooks

// F is a set of fields
type F map[string]interface{}

// Loggable allows Logger.With to consume an F.
func (f F) Loggable() map[string]interface{} { return f }

// Loggable types provide a loggable representation of their internal state.
type Loggable interface {
	Loggable() map[string]interface{}
}

// Logger provides observability
type Logger interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Trace(...interface{})
	Tracef(string, ...interface{})
	Traceln(...interface{})

	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})

	Info(...interface{})
	Infof(string, ...interface{})
	Infoln(...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnln(...interface{})

	Error(...interface{})
	Errorf(string, ...interface{})
	Errorln(...interface{})

	With(Loggable) Logger
	WithError(error) Logger
	WithField(string, interface{}) Logger
	WithFields(logrus.Fields) Logger
}

type fieldLogger struct{ log logrus.Ext1FieldLogger }

// WrapLogrus is a convenience function
func WrapLogrus(l logrus.Ext1FieldLogger) Logger {
	return fieldLogger{log: l}
}

func (l fieldLogger) Fatal(v ...interface{})              { l.log.Fatal(v...) }
func (l fieldLogger) Fatalf(fmt string, v ...interface{}) { l.log.Fatalf(fmt, v...) }
func (l fieldLogger) Fatalln(v ...interface{})            { l.log.Fatalln(v...) }
func (l fieldLogger) Trace(v ...interface{})              { l.log.Trace(v...) }
func (l fieldLogger) Tracef(fmt string, v ...interface{}) { l.log.Tracef(fmt, v...) }
func (l fieldLogger) Traceln(v ...interface{})            { l.log.Traceln(v...) }
func (l fieldLogger) Debug(v ...interface{})              { l.log.Debug(v...) }
func (l fieldLogger) Debugf(fmt string, v ...interface{}) { l.log.Debugf(fmt, v...) }
func (l fieldLogger) Debugln(v ...interface{})            { l.log.Debugln(v...) }
func (l fieldLogger) Info(v ...interface{})               { l.log.Info(v...) }
func (l fieldLogger) Infof(fmt string, v ...interface{})  { l.log.Infof(fmt, v...) }
func (l fieldLogger) Infoln(v ...interface{})             { l.log.Infoln(v...) }
func (l fieldLogger) Warn(v ...interface{})               { l.log.Warn(v...) }
func (l fieldLogger) Warnf(fmt string, v ...interface{})  { l.log.Warnf(fmt, v) }
func (l fieldLogger) Warnln(v ...interface{})             { l.log.Warnln(v...) }
func (l fieldLogger) Error(v ...interface{})              { l.log.Error(v...) }
func (l fieldLogger) Errorf(fmt string, v ...interface{}) { l.log.Errorf(fmt, v...) }
func (l fieldLogger) Errorln(v ...interface{})            { l.log.Errorln(v...) }

func (l fieldLogger) With(v Loggable) Logger {
	return l.WithFields(v.Loggable())
}

func (l fieldLogger) WithError(err error) Logger {
	return (*entry)(l.log.WithError(err))
}

func (l fieldLogger) WithField(k string, v interface{}) Logger {
	return (*entry)(l.log.WithField(k, v))
}

func (l fieldLogger) WithFields(f logrus.Fields) Logger {
	return (*entry)(l.log.WithFields(f))
}

type entry logrus.Entry

func (e *entry) Fatal(v ...interface{})              { (*logrus.Entry)(e).Fatal(v...) }
func (e *entry) Fatalf(fmt string, v ...interface{}) { (*logrus.Entry)(e).Fatalf(fmt, v...) }
func (e *entry) Fatalln(v ...interface{})            { (*logrus.Entry)(e).Fatalln(v...) }

func (e *entry) Trace(v ...interface{})              { (*logrus.Entry)(e).Trace(v...) }
func (e *entry) Tracef(fmt string, v ...interface{}) { (*logrus.Entry)(e).Tracef(fmt, v...) }
func (e *entry) Traceln(v ...interface{})            { (*logrus.Entry)(e).Traceln(v...) }

func (e *entry) Debug(v ...interface{})              { (*logrus.Entry)(e).Debug(v...) }
func (e *entry) Debugf(fmt string, v ...interface{}) { (*logrus.Entry)(e).Debugf(fmt, v...) }
func (e *entry) Debugln(v ...interface{})            { (*logrus.Entry)(e).Debugln(v...) }

func (e *entry) Info(v ...interface{})              { (*logrus.Entry)(e).Info(v...) }
func (e *entry) Infof(fmt string, v ...interface{}) { (*logrus.Entry)(e).Infof(fmt, v...) }
func (e *entry) Infoln(v ...interface{})            { (*logrus.Entry)(e).Infoln(v...) }

func (e *entry) Warn(v ...interface{})              { (*logrus.Entry)(e).Warn(v...) }
func (e *entry) Warnf(fmt string, v ...interface{}) { (*logrus.Entry)(e).Warnf(fmt, v...) }
func (e *entry) Warnln(v ...interface{})            { (*logrus.Entry)(e).Warnln(v...) }

func (e *entry) Error(v ...interface{})              { (*logrus.Entry)(e).Error(v...) }
func (e *entry) Errorf(fmt string, v ...interface{}) { (*logrus.Entry)(e).Errorf(fmt, v...) }
func (e *entry) Errorln(v ...interface{})            { (*logrus.Entry)(e).Errorln(v...) }

func (e *entry) With(v Loggable) Logger {
	return e.WithFields(v.Loggable())
}

func (e *entry) WithError(err error) Logger {
	return (*entry)((*logrus.Entry)(e).WithError(err))
}

func (e *entry) WithField(k string, v interface{}) Logger {
	return (*entry)((*logrus.Entry)(e).WithField(k, v))
}

func (e *entry) WithFields(f logrus.Fields) Logger {
	return (*entry)((*logrus.Entry)(e).WithFields(f))
}

// New logger
func New(opt ...Option) Logger {
	log := logrus.New()

	for _, option := range withDefaults(opt) {
		option(log)
	}

	return fieldLogger{log}
}
