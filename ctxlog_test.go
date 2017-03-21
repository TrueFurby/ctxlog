package ctxlog

import (
	"context"
	"testing"
)

func TestCtxLogger(t *testing.T) {
	var a bool
	if Log(nil, nil); a {
		t.Fatalf("Log error: Log method shout not have change value of a")
	}
	ctx := context.Background()
	if Log(ctx, nil); a {
		t.Fatalf("Log error: Log method shout not have change value of a")
	}
	ctx = WithLogger(ctx, func(...interface{}) {
		a = true
	})
	if logger := ctx.Value(ctxLoggerKey); logger == nil {
		t.Fatal("WithLogger Context Value error: expected Logger got nil")
	}
	if Log(ctx, nil); a != true {
		t.Fatalf("Log error: Log method did not change value of a")
	}
}

func TestCtxLoggerf(t *testing.T) {
	var a string
	if Logf(nil, "format", nil); a != "" {
		t.Fatal("Log error: Log method shout not have change value of a")
	}
	ctx := context.Background()
	if Logf(ctx, "format", nil); a != "" {
		t.Fatalf("Log error: Log method shout not have change value of a")
	}
	ctx = WithLoggerf(ctx, func(s string, v ...interface{}) {
		a = s
	})
	if logger := ctx.Value(ctxLoggerfKey); logger == nil {
		t.Fatal("WithLogger Context Value error: expected Logger got nil")
	}
	if Logf(ctx, "format"); a != "format" {
		t.Fatalf("Log error: Log method did not change value of a")
	}
}
