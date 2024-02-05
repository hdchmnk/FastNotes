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
	return nil, nil
}
