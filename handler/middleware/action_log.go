package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/TechBowl-japan/go-stations/model"
)

func ActionLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		actionLog := &model.ActionLog{
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
	})
}
