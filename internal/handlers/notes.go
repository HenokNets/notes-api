package handlers

import (
	"encoding/json"
	"net/http"
	"notes-api/internal/model"
	"sync"
	"time"
)

var (
	notes []model.Note
	mu    sync.RWMutex
)

type CreateNoteRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var req CreateNoteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	note := model.Note{
		ID:        time.Now().Format("20060102150405"),
		Title:     req.Title,
		Body:      req.Body,
		CreatedAt: time.Now(),
	}

	mu.Lock()
	notes = append(notes, note)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)
}

func ListNotes(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(notes)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	mu.RLock()
	defer mu.RUnlock()

	for _, note := range notes {
		if note.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(note)
			return
		}
	}

	http.NotFound(w, r)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var req CreateNoteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i := range notes {
		if notes[i].ID == id {
			notes[i].Title = req.Title
			notes[i].Body = req.Body

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(notes[i])

			return
		}
	}

	http.NotFound(w, r)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	mu.Lock()
	defer mu.Unlock()

	for i := range notes {
		if notes[i].ID == id {
			notes = append(notes[:i], notes[i+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}
