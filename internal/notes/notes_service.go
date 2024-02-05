package notes

import (
	"context"
	"strconv"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{repository, time.Duration(2) * time.Second}
}

func (s *service) CreateNote(c context.Context, req *CreateNoteReq) (*CreateNoteRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	n := &Note{
		UserId:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
	}

	r, err := s.Repository.CreateNote(ctx, n)
	if err != nil {
		return nil, err
	}

	res := &CreateNoteRes{
		Id:          strconv.Itoa(int(r.Id)),
		UserId:      r.UserId,
		Title:       r.Title,
		Description: r.Description,
	}
	return res, nil
}

func (s *service) GetNotesByUserID(c context.Context, id int64) (*[]Note, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	r, err := s.Repository.GetNotesByUserID(ctx, id)
	if err != nil {
		return nil, err
	}
	return r, nil
}
