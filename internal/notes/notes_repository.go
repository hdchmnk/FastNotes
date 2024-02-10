package notes

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateNote(ctx context.Context, note *Note) (*Note, error) {
	var lastInsertID int
	query := "INSERT INTO notes(userId, title, description) VALUES ($1, $2, $3) returning id"
	err := r.db.QueryRowContext(ctx, query, note.UserId, note.Title, note.Description).Scan(&lastInsertID)
	if err != nil {
		return &Note{}, err
	}
	note.Id = int64(lastInsertID)
	return note, nil
}

func (r *repository) GetNotesByUserID(ctx context.Context, id int64) (*[]Note, error) {
	query := "SELECT id, userId, title, description FROM notes WHERE userId = $1"
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var notes []Note

	for rows.Next() {
		var note Note
		if err := rows.Scan(&note.Id, &note.UserId, &note.Title, &note.Description); err != nil {
			return &notes, err
		}
		notes = append(notes, note)
	}
	if err = rows.Err(); err != nil {
		return &notes, err
	}
	return &notes, nil
}
