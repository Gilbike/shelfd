package book

import (
	"context"
	"testing"
)

func TestListTotalPages(t *testing.T) {
	type pageTestParams struct {
		returnedCount int
		expectedPages int
	}

	testParams := []pageTestParams{
		{
			returnedCount: 0,
			expectedPages: 0,
		},
		{
			returnedCount: 10,
			expectedPages: 1,
		},
		{
			returnedCount: 20,
			expectedPages: 1,
		},
		{
			returnedCount: 21,
			expectedPages: 2,
		},
	}

	for _, params := range testParams {
		mockRepo := mockRepository{
			fetchAll: func(ctx context.Context, p pagination, s sorting) ([]Book, int64, error) {
				return []Book{}, int64(params.returnedCount), nil
			},
		}

		service := NewService(&mockRepo)

		result, err := service.List(t.Context(), listFilters{})

		if err != nil {
			t.Errorf("err = %+v, expected nil, got err", err)
		}

		if result.TotalPages != params.expectedPages {
			t.Errorf("TotalPages = %d, expected %d", result.TotalPages, params.expectedPages)
		}
	}
}

func TestUploadCoverFileFormat(t *testing.T) {
	mockRepo := mockRepository{}
	service := NewService(&mockRepo)

	err := service.UploadCover(t.Context(), 0, "", []byte{}, "")

	if err == nil {
		t.Errorf("err = %+v, expected \"invalid file format\", got nil", err)
	}
}

// repository mock
type mockRepository struct {
	fetchAll func(ctx context.Context, p pagination, s sorting) ([]Book, int64, error)
	findById func(ctx context.Context, id int64) (*Book, error)
	create   func(ctx context.Context, book *Book) (int64, error)
	delete   func(ctx context.Context, id int64) error
	update   func(ctx context.Context, id int64, dirtyFields map[string]any) error
}

func (mr *mockRepository) FetchAll(ctx context.Context, p pagination, s sorting) ([]Book, int64, error) {
	return mr.fetchAll(ctx, p, s)
}
func (mr *mockRepository) FindById(ctx context.Context, id int64) (*Book, error) {
	return mr.findById(ctx, id)
}
func (mr *mockRepository) Create(ctx context.Context, book *Book) (int64, error) {
	return mr.create(ctx, book)
}
func (mr *mockRepository) Delete(ctx context.Context, id int64) error {
	return mr.delete(ctx, id)
}
func (mr *mockRepository) Update(ctx context.Context, id int64, dirtyFields map[string]any) error {
	return mr.update(ctx, id, dirtyFields)
}
