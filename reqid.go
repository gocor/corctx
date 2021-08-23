package corctx

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type ctxKey int

const reqIDKey ctxKey = ctxKey(0)

// GetRequestID retuns the requestID if found from context
func GetRequestID(ctx context.Context) (string, bool) {
	rid, ok := ctx.Value(reqIDKey).(string)
	return rid, ok
}

// WithRequestID will give a new context with a requestID if one does not exist on the context
func WithRequestID(ctx context.Context, requestID string) context.Context {
	// do not overwrite a request id if it is already on context
	if _, ok := GetRequestID(ctx); ok {
		return ctx
	}
	return withRequestID(ctx, requestID)
}

func withRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, reqIDKey, requestID)
}

// WithHTTPRequest will return a context with RequestID set from http x-request-id header
// or if it does not exist, it will create one
func WithHTTPRequest(r *http.Request) context.Context {
	reqID := r.Header.Get("X-Request-ID")
	if len(reqID) == 0 {
		reqID = uuid.New().String()
	}
	return WithRequestID(r.Context(), reqID)
}
