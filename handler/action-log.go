package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/TechBowl-japan/go-stations/model"
)

// A ActionLogHandler implements ActionLog endpoint.
type ActionLogHandler struct{}

// NewActionLogHandler returns ActionLogHandler based http.Handler.
func NewActionLogHandler() *ActionLogHandler {
	return &ActionLogHandler{}
}

// ServeHTTP implements http.Handler interface.
func (a *ActionLogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(50 * time.Millisecond)
	response := &model.ActionLogResponse{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}
