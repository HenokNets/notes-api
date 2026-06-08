package server

import (
	"net/http"
	"notes-api/internal/handlers"

	_ "modernc.org/sqlite"
)

func NewMux(
	notesHandler *handlers.NotesHandler,
) *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc(
		"POST /notes",
		notesHandler.CreateNote,
	)

	mux.HandleFunc(
		"GET /notes",
		notesHandler.ListNotes,
	)

	mux.HandleFunc(
		"GET /notes/{id}",
		notesHandler.GetNote,
	)

	mux.HandleFunc(
		"PUT /notes/{id}",
		notesHandler.UpdateNote,
	)

	mux.HandleFunc(
		"DELETE /notes/{id}",
		notesHandler.DeleteNote,
	)

	return mux

}
