package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrGetComment     = errors.New("error in getting comment by id")
	ErrUpdateComment  = errors.New("error in updating a comment")
	ErrDeleteComment  = errors.New("error in deleting a comment")
	ErrGetAllComments = errors.New("error in getting all comments")
	ErrCreateComment  = errors.New("error in creating all comments")
)

// Comment - description struct of comment
type Comment struct {
	Id     string
	Slug   string
	Body   string
	Author string
}

// Store - defines all methods for our service
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	PostComment(context.Context, Comment) (Comment, error)
}

// Service - struct for logic
type Service struct {
	Store Store
}

// NewService - returns a new service
func NewService(store Store) *Service {
	return &Service{Store: store}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	comment, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrGetComment
	}
	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, id string, cmt Comment) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, id, cmt)
	if err != nil {
		return Comment{}, fmt.Errorf("error in updating comment")
	}
	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}

func (s *Service) PostComment(ctx context.Context, comment Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, comment)
	if err != nil {
		return Comment{}, err
	}
	return insertedCmt, nil
}

