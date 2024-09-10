package middlewares

import (
	"context"
	"sync"
)

type traceIDKey struct {}

var (
	logNo int = 1
	mu sync.Mutex
)

func newTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	logNo += 1
	mu.Unlock()

	return no
}

func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func GetTraceID(ctx context.Context) int {
	id := ctx.Value(traceIDKey{})

	if idInt, ok := id.(int); ok {
		return idInt
	}

	return 0
}