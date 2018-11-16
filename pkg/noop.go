package log

import "github.com/sirupsen/logrus"

type noop struct{}

func noState(error)                               {}
func (noop) Fatal(v ...interface{})               { logrus.Fatal(v...) }
func (noop) Fatalf(fmt string, v ...interface{})  { logrus.Fatalf(fmt, v...) }
func (noop) Fatalln(v ...interface{})             { logrus.Fatalln(v...) }
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
func (noop) WithLocus(string) Logger              { return noop{} }
func (noop) WithError(error) Logger               { return noop{} }
func (noop) WithField(string, interface{}) Logger { return noop{} }
func (noop) WithFields(F) Logger                  { return noop{} }
func (noop) State(func(Logger)) State             { return state(func() {}) }
func (noop) IfErr(func(Logger)) ErrState          { return errState(noState) }
func (noop) IfNoErr(func(Logger)) ErrState        { return errState(noState) }
