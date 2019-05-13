package log

import "github.com/sirupsen/logrus"

type noop struct{}

func noState(error)                               {}
func (noop) Fatal(v ...interface{})               { logrus.Fatal(v...) }
func (noop) Fatalf(fmt string, v ...interface{})  { logrus.Fatalf(fmt, v...) }
func (noop) Fatalln(v ...interface{})             { logrus.Fatalln(v...) }
func (noop) Trace(...interface{})                 {}
func (noop) Tracef(string, ...interface{})        {}
func (noop) Traceln(...interface{})               {}
func (noop) Debug(...interface{})                 {}
func (noop) Debugf(string, ...interface{})        {}
func (noop) Debugln(...interface{})               {}
func (noop) Info(...interface{})                  {}
func (noop) Infof(string, ...interface{})         {}
func (noop) Infoln(...interface{})                {}
func (noop) Warn(...interface{})                  {}
func (noop) Warnf(string, ...interface{})         {}
func (noop) Warnln(...interface{})                {}
func (noop) Error(...interface{})                 {}
func (noop) Errorf(string, ...interface{})        {}
func (noop) Errorln(...interface{})               {}
func (noop) WithError(error) Logger               { return noop{} }
func (noop) WithField(string, interface{}) Logger { return noop{} }
func (noop) WithFields(F) Logger                  { return noop{} }
