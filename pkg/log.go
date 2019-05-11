package log

import (
	"unsafe"

	"github.com/lthibault/logrus"
)

const (
	locusLabel = "locus"
)

// F is a set of fields
type F map[string]interface{}

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

	WithError(error) Logger
	WithField(string, interface{}) Logger
	WithFields(F) Logger
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

func (l fieldLogger) WithError(err error) Logger {
	return (*entry)(unsafe.Pointer(l.log.WithError(err)))
}

func (l fieldLogger) WithField(k string, v interface{}) Logger {
	return (*entry)(unsafe.Pointer(l.log.WithField(k, v)))
}

func (l fieldLogger) WithFields(f F) Logger {
	return (*entry)(unsafe.Pointer(l.log.WithFields(logrus.Fields(f))))
}

type entry logrus.Entry

func (e *entry) Fatal(v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Fatal(v...)
}
func (e *entry) Fatalf(fmt string, v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Fatalf(fmt, v...)
}
func (e *entry) Fatalln(v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Fatalln(v...)
}

func (e *entry) Trace(v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Trace(v...)
}
func (e *entry) Tracef(fmt string, v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Tracef(fmt, v...)
}
func (e *entry) Traceln(v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Traceln(v...)
}

func (e *entry) Debug(v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Debug(v...)
}
func (e *entry) Debugf(fmt string, v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Debugf(fmt, v...)
}
func (e *entry) Debugln(v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Debugln(v...)
}

func (e *entry) Info(v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Info(v...)
}
func (e *entry) Infof(fmt string, v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Infof(fmt, v...)
}
func (e *entry) Infoln(v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Infoln(v...)
}

func (e *entry) Warn(v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Warn(v...)
}
func (e *entry) Warnf(fmt string, v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Warnf(fmt, v...)
}
func (e *entry) Warnln(v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Warnln(v...)
}

func (e *entry) Error(v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Error(v...)
}
func (e *entry) Errorf(fmt string, v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Errorf(fmt, v...)
}
func (e *entry) Errorln(v ...interface{}) {
	(*logrus.Entry)(unsafe.Pointer(e)).Errorln(v...)
}

func (e *entry) WithError(err error) Logger {
	return (*entry)(unsafe.Pointer(
		(*logrus.Entry)(unsafe.Pointer(e)).WithError(err),
	))
}
func (e *entry) WithField(k string, v interface{}) Logger {
	return (*entry)(unsafe.Pointer(
		(*logrus.Entry)(unsafe.Pointer(e)).WithField(k, v),
	))
}
func (e *entry) WithFields(f F) Logger {
	return (*entry)(unsafe.Pointer(
		(*logrus.Entry)(unsafe.Pointer(e)).WithFields(
			logrus.Fields(f),
		),
	))
}

// New logger
func New(opt ...Option) Logger {
	var s spec

	// Defaults
	OptLevel(InfoLevel)(&s)

	for _, fn := range opt {
		fn(&s)
	}
	return s.mkLogger()
}
