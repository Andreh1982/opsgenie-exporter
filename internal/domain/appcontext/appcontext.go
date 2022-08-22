package appcontext

import (
	"context"
)

// ContextKey is the key for the app context
type ContextKey string

const (
	// AppContextKey is the key for the app context
	AppContextKey               ContextKey = "appContext"
	defaultBackgroundContextKey ContextKey = "ctx"
)

// Context is wrapper for gin context 'n go context. Provide logger with trace_id.
type Context interface {
	Context() context.Context
}

// New returns a new app context
func New(ctx context.Context) Context {
	return &appContext{
		defaultBackgroundContext: ctx,
	}
}

// NewBackground returns a new background void context
func NewBackground() Context {
	ctx := context.Background()

	return &appContext{
		defaultBackgroundContext: ctx,
	}
}

type appContext struct {
	defaultBackgroundContext context.Context
}

func (appContext *appContext) Context() context.Context {
	return appContext.defaultBackgroundContext
}
