package main

import (
	"database/sql"
	"log"
	"net/http"
	"notes-api/internal/handlers"
	"notes-api/internal/server"
	"notes-api/internal/store"
	"os"
)

func main() {
	db, err := sql.Open(
		"sqlite",
		"notes.db",
	)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	schema, err := os.ReadFile(
		"schema.sql",
	)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(
		string(schema),
	)

	if err != nil {
		log.Fatal(err)
	}

	store := store.New(db)

	notesHandler := handlers.NewNotesHandler(
		store,
	)

	mux := server.NewMux(
		notesHandler,
	)

	log.Println("server running on :8080")

	err = http.ListenAndServe(
		":8080",
		mux,
	)

	if err != nil {
		log.Fatal(err)
	}
}
