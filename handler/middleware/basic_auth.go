package middleware

import (
	"net/http"
	"os"
)

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := os.Getenv("BASIC_AUTH_USER_ID")
		password := os.Getenv("BASIC_AUTH_PASSWORD")
		if userID == "" || password == "" {
			http.Error(w, "basic auth is required", http.StatusUnauthorized)
			return
		}

		user, pass, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "basic auth is required", http.StatusUnauthorized)
			return
		}
		if user != userID || pass != password {
			http.Error(w, "user or password is incorrect", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
