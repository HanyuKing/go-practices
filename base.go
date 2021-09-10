package main

import "context"

func main() {
	SpanFromContext(context.TODO())
}

func SpanFromContext(ctx context.Context) Span {
	val := ctx.Value(activeSpanKey)
	if sp, ok := val.(Span); ok {
		return sp
	}
	return nil
}

type Span interface {
}

type contextKey struct{}

var activeSpanKey = contextKey{}
