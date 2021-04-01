// Package zaputil contains utilities for working with go.uber.org/zap.
package zaputil

import (
	"github.com/lthibault/log"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

type logWrapper zap.SugaredLogger

// Wrap a zap logger in a shim satisfying github.com/lthibault/log.Logger.
//
// Trace messages are ignored.  Expect degraded performance.
func Wrap(log *zap.SugaredLogger) log.Logger { return (*logWrapper)(log) }

func (w *logWrapper) zap() *zap.SugaredLogger { return (*zap.SugaredLogger)(w) }

func (w *logWrapper) Fatal(v ...interface{})              { w.zap().Fatal(v...) }
func (w *logWrapper) Fatalf(fmt string, v ...interface{}) { w.zap().Fatalf(fmt, v...) }
func (w *logWrapper) Fatalln(v ...interface{})            { w.Fatal(v...) }

func (w *logWrapper) Trace(...interface{})          {}
func (w *logWrapper) Tracef(string, ...interface{}) {}
func (w *logWrapper) Traceln(...interface{})        {}

func (w *logWrapper) Debug(v ...interface{})              { w.zap().Debug(v...) }
func (w *logWrapper) Debugf(fmt string, v ...interface{}) { w.zap().Debugf(fmt, v...) }
func (w *logWrapper) Debugln(v ...interface{})            { w.Debug(v...) }

func (w *logWrapper) Info(v ...interface{})              { w.zap().Info(v...) }
func (w *logWrapper) Infof(fmt string, v ...interface{}) { w.zap().Infof(fmt, v...) }
func (w *logWrapper) Infoln(v ...interface{})            { w.Info(v...) }

func (w *logWrapper) Warn(v ...interface{})              { w.zap().Warn(v...) }
func (w *logWrapper) Warnf(fmt string, v ...interface{}) { w.zap().Warnf(fmt, v...) }
func (w *logWrapper) Warnln(v ...interface{})            { w.Warn(v...) }

func (w *logWrapper) Error(v ...interface{})              { w.zap().Error(v...) }
func (w *logWrapper) Errorf(fmt string, v ...interface{}) { w.zap().Errorf(fmt, v...) }
func (w *logWrapper) Errorln(v ...interface{})            { w.Error(v...) }

func (w *logWrapper) With(l log.Loggable) log.Logger {
	return w.WithFields(l.Loggable())
}

func (w *logWrapper) WithError(err error) log.Logger {
	return Wrap(w.zap().Desugar().With(zap.Error(err)).Sugar())
}

func (w *logWrapper) WithField(name string, v interface{}) log.Logger {
	return Wrap(w.zap().With(name, v))
}

func (w *logWrapper) WithFields(f logrus.Fields) log.Logger {
	args := make([]interface{}, 0, len(f)*2)
	for key, value := range f {
		args = append(args, key)
		args = append(args, value)
	}

	return Wrap(w.zap().With(args...))
}
