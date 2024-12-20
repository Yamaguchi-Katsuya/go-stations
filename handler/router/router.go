package router

import (
	"database/sql"
	"net/http"

	"github.com/TechBowl-japan/go-stations/handler"
	"github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/service"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	mux.Handle("/healthz", handler.NewHealthzHandler())
	mux.Handle("/basic-auth", middleware.BasicAuth(handler.NewHealthzHandler()))
	mux.Handle("/todos", handler.NewTODOHandler(service.NewTODOService(todoDB)))
	mux.Handle("/do-panic", handler.NewDoPanicHandler())
	mux.Handle("/recovery-panic", middleware.Recovery(handler.NewDoPanicHandler()))
	mux.Handle("/os", middleware.OS(handler.NewOSHandler()))
	mux.Handle("/action-log", middleware.OS(middleware.ActionLog(handler.NewActionLogHandler())))
	mux.Handle("/sleep", handler.NewSleepHandler())
	return mux
}
