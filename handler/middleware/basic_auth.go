package middleware

import (
	"log"
	"net/http"
	"os"
)

var (
	basicAuthUserID   string
	basicAuthPassword string
)

func init() {
	basicAuthUserID = os.Getenv("BASIC_AUTH_USER_ID")
	basicAuthPassword = os.Getenv("BASIC_AUTH_PASSWORD")

	if basicAuthUserID == "" || basicAuthPassword == "" {
		log.Fatal("BASIC_AUTH_USER_ID or BASIC_AUTH_PASSWORD is not set")
	}
}

func BasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
			http.Error(w, "basic auth is required", http.StatusUnauthorized)
			return
		}
		if user != basicAuthUserID || pass != basicAuthPassword {
			http.Error(w, "user or password is incorrect", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
