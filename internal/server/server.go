package server

import (
	"net/http"
	"notes-api/internal/handlers"
)

func NewServer() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.HealthHandler)
	mux.HandleFunc("GET /version", handlers.VersionHandler)
	mux.HandleFunc("POST /notes", handlers.CreateNote)

	mux.HandleFunc("GET /notes", handlers.ListNotes)

	mux.HandleFunc("GET /notes/{id}", handlers.GetNote)

	mux.HandleFunc("PUT /notes/{id}", handlers.UpdateNote)

	mux.HandleFunc("DELETE /notes/{id}", handlers.DeleteNote)
	return mux
}
