package log

import "github.com/sirupsen/logrus"

func Fatal(v ...interface{})              { logrus.StandardLogger().Fatal(v...) }
func Fatalf(fmt string, v ...interface{}) { logrus.StandardLogger().Fatalf(fmt, v...) }
func Fatalln(v ...interface{})            { logrus.StandardLogger().Fatalln(v...) }
func Trace(v ...interface{})              { logrus.StandardLogger().Trace(v...) }
func Tracef(fmt string, v ...interface{}) { logrus.StandardLogger().Tracef(fmt, v...) }
func Traceln(v ...interface{})            { logrus.StandardLogger().Traceln(v...) }
func Debug(v ...interface{})              { logrus.StandardLogger().Debug(v...) }
func Debugf(fmt string, v ...interface{}) { logrus.StandardLogger().Debugf(fmt, v...) }
func Debugln(v ...interface{})            { logrus.StandardLogger().Debugln(v...) }
func Info(v ...interface{})               { logrus.StandardLogger().Info(v...) }
func Infof(fmt string, v ...interface{})  { logrus.StandardLogger().Infof(fmt, v...) }
func Infoln(v ...interface{})             { logrus.StandardLogger().Infoln(v...) }
func Warn(v ...interface{})               { logrus.StandardLogger().Warn(v...) }
func Warnf(fmt string, v ...interface{})  { logrus.StandardLogger().Warnf(fmt, v...) }
func Warnln(v ...interface{})             { logrus.StandardLogger().Warnln(v...) }
func Error(v ...interface{})              { logrus.StandardLogger().Error(v...) }
func Errorf(fmt string, v ...interface{}) { logrus.StandardLogger().Errorf(fmt, v...) }
func Errorln(v ...interface{})            { logrus.StandardLogger().Errorln(v...) }
