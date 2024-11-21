package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/model"
)

// A OSHandler implements OS endpoint.
type OSHandler struct{}

// NewOSHandler returns OSHandler based http.Handler.
func NewOSHandler() *OSHandler {
	return &OSHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *OSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	os := middleware.GetOS(r.Context())
	response := &model.OSResponse{
		OS: os,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}
}
