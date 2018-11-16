package log

import "context"

// Get Logger
func Get(c context.Context) Logger {
	return c.Value(keyLogger).(Logger)
}

// Set logger
func Set(c context.Context, l Logger) context.Context {
	return context.WithValue(c, keyLogger, l)
}

// Maybe returns the logger, defaulting to a NoOp
func Maybe(c context.Context) Logger {
	if l, ok := c.Value(keyLogger).(Logger); ok {
		return l
	}
	return noop{}
}
