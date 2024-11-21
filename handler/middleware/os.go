package middleware

import (
	"context"
	"net/http"
)

type ctxKey struct{}

func OS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ua := r.UserAgent()
		if ua == "" {
			http.Error(w, "user agent is required", http.StatusBadRequest)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, ctxKey{}, ua)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetOS(ctx context.Context) string {
	if os, ok := ctx.Value(ctxKey{}).(string); ok {
		return os
	}
	return ""
}
