package server

import (
	"net/http"
	"notes-api/internal/handlers"
)

func NewServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.HealthHandler)
	mux.HandleFunc("GET /version", handlers.VersionHandler)
	return mux
}
