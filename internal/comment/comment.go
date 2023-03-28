package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrGetComment     = errors.New("Error in getting comment by id")
	ErrUpdateComment  = errors.New("Error in updating a comment")
	ErrDeleteComment  = errors.New("Error in deleting a comment")
	ErrGetAllComments = errors.New("Error in getting all comments")
	ErrCreateComment  = errors.New("Error in creating all comments")
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
	UpdateComment(context.Context, string) (Comment, error)
	DeleteComment(ctx context.Context, id string) error
	CreateComment(ctx context.Context, comment Comment) (Comment, error)
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

func (s *Service) UpdateComment(ctx context.Context, id string) error {
	fmt.Println("Updating a comment")
	_, err := s.Store.UpdateComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return ErrUpdateComment
	}
	return nil
}

func (S *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrDeleteComment
}

func (s *Service) CreateComment(ctx context.Context, comment Comment) (Comment, error) {
	return Comment{}, ErrCreateComment
}

//func (s *Service) GetAllComments(ctx context.Context) (Comment,error) {
//	return Comment{},ErrGetAllComments
//}


