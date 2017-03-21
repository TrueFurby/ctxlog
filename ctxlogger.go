package ctxlog

import "context"

type key int

const ctxLoggerKey key = 0

//Logger defines the type that must be met in order to use the WithLogger context function.
type Logger func(v ...interface{})

//WithLogger returns a copy of parent in which the Logger method is associtated with the non exportable ctxLoggerKey.
func WithLogger(parent context.Context, logger Logger) context.Context {
	return context.WithValue(parent, ctxLoggerKey, logger)
}

//Log takes a context and will use the Logger method associated with the non exported cxtLoggerKey value.
func Log(ctx context.Context, v ...interface{}) {
	if ctx == nil {
		return
	}
	logger, ok := ctx.Value(ctxLoggerKey).(Logger)
	if !ok {
		return
	}
	logger(v...)
}
