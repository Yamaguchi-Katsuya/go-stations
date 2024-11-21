package handler

import (
	"net/http"
)

// A DoPanicHandler implements panic endpoint.
type DoPanicHandler struct{}

// NewDoPanicHandler returns DoPanicHandler based http.Handler.
func NewDoPanicHandler() *DoPanicHandler {
	return &DoPanicHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *DoPanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	panic("panic")
}
