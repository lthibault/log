package log

import "github.com/sirupsen/logrus"

// Fatal prints each value then calls os.Exit(1)
func Fatal(v ...interface{}) { logrus.StandardLogger().Fatal(v...) }

// Fatalf provides fmt.Printf-like formatting for a call to Fatal.
func Fatalf(fmt string, v ...interface{}) { logrus.StandardLogger().Fatalf(fmt, v...) }

// Fatalln calls Fatal, appending a newline to the log output.
func Fatalln(v ...interface{}) { logrus.StandardLogger().Fatalln(v...) }

// Trace prints a message at trace level.
func Trace(v ...interface{}) { logrus.StandardLogger().Trace(v...) }

// Tracef provides fmt.Printf-like formatting for a call to Trace
func Tracef(fmt string, v ...interface{}) { logrus.StandardLogger().Tracef(fmt, v...) }

// Traceln calls Trace, adding a newline to the log output.
func Traceln(v ...interface{}) { logrus.StandardLogger().Traceln(v...) }

// Debug prints a message at debug level.
func Debug(v ...interface{}) { logrus.StandardLogger().Debug(v...) }

// Debugf provides fmt.Printf-like formatting for a call to Trace
func Debugf(fmt string, v ...interface{}) { logrus.StandardLogger().Debugf(fmt, v...) }

// Debugln calls Trace, adding a newline to the log output.
func Debugln(v ...interface{}) { logrus.StandardLogger().Debugln(v...) }

// Info prints a message at info level.
func Info(v ...interface{}) { logrus.StandardLogger().Info(v...) }

// Infof provides fmt.Printf-like formatting for a call to Trace
func Infof(fmt string, v ...interface{}) { logrus.StandardLogger().Infof(fmt, v...) }

// Infoln calls Trace, adding a newline to the log output.
func Infoln(v ...interface{}) { logrus.StandardLogger().Infoln(v...) }

// Warn prints a message at warn level.
func Warn(v ...interface{}) { logrus.StandardLogger().Warn(v...) }

// Warnf provides fmt.Printf-like formatting for a call to Trace
func Warnf(fmt string, v ...interface{}) { logrus.StandardLogger().Warnf(fmt, v...) }

// Warnln calls Trace, adding a newline to the log output.
func Warnln(v ...interface{}) { logrus.StandardLogger().Warnln(v...) }

// Error prints a message at error level.
func Error(v ...interface{}) { logrus.StandardLogger().Error(v...) }

// Errorf provides fmt.Printf-like formatting for a call to Trace
func Errorf(fmt string, v ...interface{}) { logrus.StandardLogger().Errorf(fmt, v...) }

// Errorln calls Trace, adding a newline to the log output.
func Errorln(v ...interface{}) { logrus.StandardLogger().Errorln(v...) }
