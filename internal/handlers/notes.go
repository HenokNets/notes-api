package handlers

import (
	"encoding/json"
	"net/http"
	"notes-api/internal/model"
	"notes-api/internal/store"
	"time"
)

type NotesHandler struct {
	store *store.Store
}

func NewNotesHandler(s *store.Store) *NotesHandler {
	return &NotesHandler{
		store: s,
	}
}

type CreateNoteRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (h *NotesHandler) CreateNote(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req CreateNoteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	now := time.Now()

	note := model.Note{
		ID:        now.Format("20060102150405"),
		UserID:    "anonymous",
		Title:     req.Title,
		Body:      req.Body,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := h.store.CreateNote(note); err != nil {
		http.Error(
			w,
			"failed to create note",
			http.StatusInternalServerError
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application,json"
	)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(note)
}

func (h *NotesHandler) ListNotes(
	w http.ResponseWriter,
	r *http.Request,
) {
	notes, err := h.store.ListNotes()

	if err != nil {
		http.Error(
			w, 
			"database error",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(notes)
}

func (h *NotesHandler) GetNote(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := r.PathValue("id")

	note, err := h.store.GetNote(id)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json"
	)

	json.NewEncoder(w).Encode(note)
}

func (h *NotesHandler) UpdateNote(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := r.PathValue("id")

	var req CreateNoteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(
			w,
			"invalid json",
			http.StatusBadRequest
		)
		return
	}

	err := h.store.UpdateNote(
		id,
		req.Title,
		req.Body,
	)

	if err != nil {
		http.Error(
			w,
			"database error",
			http.StatusInternalServerError,
		)
		return
	}

	note, err := h.store.GetNote(id)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(note)
}

func (h *NotesHandler) DeleteNote(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := r.PathValue("id")

	err := h.store.DeleteNote(id) 

	if err != nil {
		http.Error(
			w,
			"database error",
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}