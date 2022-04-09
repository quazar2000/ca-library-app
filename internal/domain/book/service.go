package book

import (
	"context"

	"ca-library-app/internal/adapters/api/book"

	"github.com/quazar2000/ca-library-app/internal/adapters/api/author"
)

type service struct {
	storage       Storage
	authorService author.Service
	genreService  genre.Service
}

func NewService(storage Storage) book.Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, dto *CreateBookDTO) *Book {
	author := s.authorService.GetByUUID(ctx, dto.AuthorUUID)
	genre := s.genreService.GetByUUID(ctx, dto.GenreUUID)
	return nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) *Book {
	return s.storage.GetOne(uuid)
}

func (s *service) GetAll(ctx context.Context, liimt, offset int) []*Book {
	return s.storage.GetAll(limit, offset)
}
