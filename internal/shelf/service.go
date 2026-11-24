package shelf

import "context"

type repository interface {
	FetchBookShelfForUser(ctx context.Context, userId int64, bookId int64) (string, error)
	FetchBookUserCount(ctx context.Context, bookId int64) (int64, error)
	UpsertShelf(ctx context.Context, userId int64, bookId int64, collection string) error
}

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	return &Service{repository}
}

func (s *Service) SetBookShelf(ctx context.Context, payload *UserBook) error {
	if err := payload.IsCollectionNameValid(); err != nil {
		return err
	}
	return s.repository.UpsertShelf(ctx, payload.UserId, payload.BookId, payload.ShelfName)
}

func (s *Service) GetBookCollection(ctx context.Context, userId int64, bookId int64) *string {
	collection, err := s.repository.FetchBookShelfForUser(ctx, userId, bookId)
	if err != nil {
		return nil
	}
	return &collection
}

func (s *Service) IsBookInCollection(ctx context.Context, bookId int64) (bool, error) {
	count, err := s.repository.FetchBookUserCount(ctx, bookId)
	if err != nil {
		return false, err
	}
	return count != 0, err
}
