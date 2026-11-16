package collection

import "context"

type repository interface {
	UpsertBuiltIn(ctx context.Context, userId int64, bookId int64, collection string) error
}

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	return &Service{repository}
}

func (s *Service) SetBookCollection(ctx context.Context, payload *UserBook) error {
	if err := payload.IsCollectionNameValid(); err != nil {
		return err
	}
	return s.repository.UpsertBuiltIn(ctx, payload.UserId, payload.BookId, payload.CollectionName)
}
