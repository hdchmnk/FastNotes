package notes

import "context"

type Note struct {
	Id          int64  `json:"id" db:"id"`
	UserId      int64  `json:"userId" db:"userId"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type CreateNoteReq struct {
	UserId      int64  `json:"userId" db:"userId"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type CreateNoteRes struct {
	Id          string `json:"id" db:"id"`
	UserId      int64  `json:"userId" db:"userId"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type Repository interface {
	CreateNote(ctx context.Context, note *Note) (*Note, error)
	GetNotesByUserID(ctx context.Context, id int64) (*[]Note, error)
}

type Service interface {
	CreateNote(c context.Context, req *CreateNoteReq) (*CreateNoteRes, error)
	GetNotesByUserID(ctx context.Context, id int64) (*[]Note, error)
}
