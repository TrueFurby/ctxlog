package ctxlog

import "context"

const ctxLoggerfKey key = 1

//Loggerf defines the type that must be met in order to use the WithLoggerf context function.
type Loggerf func(format string, v ...interface{})

//WithLoggerf returns a copy of parent in which the Loggerf method is associtated with the non exportable ctxLoggerfKey.
func WithLoggerf(parent context.Context, logger Loggerf) context.Context {
	return context.WithValue(parent, ctxLoggerfKey, logger)
}

//Logf takes a context and will use the Loggerf method associated with the non exported cxtLoggerfKey value.
func Logf(ctx context.Context, format string, v ...interface{}) {
	if ctx == nil {
		return
	}
	logger, ok := ctx.Value(ctxLoggerfKey).(Loggerf)
	if !ok {
		return
	}
	logger(format, v...)
}
