package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type actionLog struct {
	Timestamp time.Time `json:"timestamp"`
	Latency   int64     `json:"latency"`
	Path      string    `json:"path"`
	OS        string    `json:"os"`
}

func ActionLog(next http.Handler) http.Handler {
	return OS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		actionLog := &actionLog{
			Timestamp: time.Now(),
			Path:      r.URL.Path,
			OS:        GetOS(r.Context()),
		}
		defer func() {
			actionLog.Latency = time.Since(actionLog.Timestamp).Milliseconds()
			json.NewEncoder(w).Encode(actionLog)
			fmt.Println(actionLog)
		}()
		next.ServeHTTP(w, r)
	}))
}
