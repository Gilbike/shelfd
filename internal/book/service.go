package book

import (
	"context"
)

const pageSize = 20

type repository interface {
	FetchAll(ctx context.Context, p pagination, s sorting) ([]Book, int64, error)
}

type listFilters struct {
	page   int
	sortBy string
}

type listResult struct {
	Books       []Book
	TotalCount  int64
	TotalPages  int
	CurrentPage int
}

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) List(ctx context.Context, filters listFilters) (*listResult, error) {
	books, bookCount, err := s.repository.FetchAll(ctx, pagination{page: filters.page, pageSize: pageSize}, sorting{by: filters.sortBy, descending: false})
	if err != nil {
		return nil, err
	}

	return &listResult{
		Books:       books,
		TotalCount:  bookCount,
		CurrentPage: filters.page,
		TotalPages:  (int(bookCount) + pageSize - 1) / pageSize,
	}, nil
}
