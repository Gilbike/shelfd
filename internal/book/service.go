package book

import (
	"context"
	"time"
)

const pageSize = 20

type repository interface {
	FetchAll(ctx context.Context, p pagination, s sorting) ([]Book, int64, error)
	Create(ctx context.Context, book *Book) (int64, error)
}

type listFilters struct {
	page   int
	sortBy string
}

type listResult struct {
	Books       []Book
	PageSize    int
	TotalCount  int64
	TotalPages  int
	CurrentPage int
}

type createPayload struct {
	title         string
	authors       []string
	pages         int
	isbn          *string
	publishedYear *int
	description   *string
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
		PageSize:    pageSize,
		TotalCount:  bookCount,
		CurrentPage: filters.page,
		TotalPages:  (int(bookCount) + pageSize - 1) / pageSize,
	}, nil
}

func (s *Service) Create(ctx context.Context, payload createPayload) (*Book, error) {
	book := &Book{
		Title:         payload.title,
		Authors:       payload.authors,
		Pages:         payload.pages,
		ISBN:          payload.isbn,
		PublishedYear: payload.publishedYear,
		Description:   payload.description,
	}

	insertId, err := s.repository.Create(ctx, book)
	if err != nil {
		return nil, err
	}

	book.ID = insertId

	now := time.Now()
	book.CreatedAt = now
	book.UpdatedAt = now

	return book, nil
}
