package store

import (
	"database/sql"
	"notes-api/internal/model"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateNote(note model.Note) error {
	_, err := s.db.Exec(
		`
		INSERT INTO notes (
			id, 
			user_id,
			title,
			body,
			created_at,
			updated_at
		)
		VALUES (?, ?, ?, ?, ?, ?)
		`,
		note.ID,
		note.UserID,
		note.Title,
		note.Body,
		note.CreatedAt,
		note.UpdatedAt,
	)

	return err
}

func (s *Store) ListNotes() ([]model.Note, error) {

	rows, err := s.db.Query(
		`
		SELECT
			id,
			user_id,
			title, 
			body,
			created_at,
			updated_at,
		FROM notes
		ORDER BY created_at DESC
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var notes []model.Note

	for rows.Next() {
		var note model.Note

		err := rows.Scan(
			&note.ID,
			&note.UserID,
			&note.Title,
			&note.Body,
			&note.CreatedAt,
			&note.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}

func (s *Store) GetNote(id string) (model.Note, error) {

	var note model.Note

	err := s.db.QueryRow(
		`
		SELECT
			id,
			user_id,
			title,
			body,
			created_at,
			updated_at,
		FROM notes
		WHERE id = ?
		`,
		id,
	).Scan(
		&note.ID,
		&note.UserID,
		&note.Title,
		&note.Body,
		&note.CreatedAt,
		&note.UpdatedAt,
	)

	return note, err
}

func (s *Store) UpdateNote(id, title, body string) error {
	_, err := s.db.Exec(
		`
		UPDATE notes 
		SET 
			title = ?,
			body = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
		`,
		title,
		body,
		id,
	)

	return err
}

func (s *Store) DeleteNote(id string) error {

	_, err := s.db.Exec(
		`
		DELETE FROM notes
		WHERE id = ?
		`,
		id,
	)

	return err
}
