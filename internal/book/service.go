package book

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"
)

const pageSize = 20

type repository interface {
	FetchAll(ctx context.Context, p pagination, s sorting) ([]Book, int64, error)
	FindById(ctx context.Context, id int64) (*Book, error)
	Create(ctx context.Context, book *Book) (int64, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, dirtyFields map[string]any) error
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

func (s *Service) Get(ctx context.Context, id int64) (*Book, error) {
	return s.repository.FindById(ctx, id)
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

	err := book.Validate()
	if err != nil {
		return nil, err
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

func (s *Service) UploadCover(ctx context.Context, id int64, mimeType string, content []byte, urlTemplate string) error {
	parts := strings.Split(mimeType, "/")
	if len(parts) < 2 {
		return fmt.Errorf("invalid file format")
	}

	var fileExtension string
	// svg mime type is svg+xml
	// https://developer.mozilla.org/en-US/docs/Web/Media/Guides/Formats/Image_types#common_image_file_types
	if strings.Contains(parts[1], "+") {
		fileExtension = "svg"
	} else {
		fileExtension = parts[1]
	}

	fileName := fmt.Sprintf("b%d.%s", id, fileExtension)

	// TODO: move to data folder with db
	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		return fmt.Errorf("failed to write cover file: %w", err)
	}

	url := fmt.Sprintf(urlTemplate, fileName)

	// TODO: change to more typesafe implementation
	err = s.repository.Update(ctx, id, map[string]any{"cover_url": url})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
