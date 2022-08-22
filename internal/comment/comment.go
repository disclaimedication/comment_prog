package comment

import (
	"context"
	"fmt"
)

//This package holds the comment structure
//and interactive functions for the comment

type comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

//It is an interface between the database
//and the comment function
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

//constructor function for defining various
//methods in the comment function
type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("retrieving the comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return cmt, nil
}
