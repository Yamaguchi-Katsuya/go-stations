package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/TechBowl-japan/go-stations/model"
)

// A SleepHandler implements sleep endpoint.
type SleepHandler struct{}

// NewSleepHandler returns SleepHandler based http.Handler.
func NewSleepHandler() *SleepHandler {
	return &SleepHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *SleepHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	response := &model.HealthzResponse{
		Message: "sleep action is done",
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}
