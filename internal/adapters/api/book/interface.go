package book

import (
	"context"

	"github.com/quazar2000/ca-library-app/internal/domain/book"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *book.Book
	GetAll(ctx context.Context, limit, offset int) []*book.Book
	Create(ctx context.Context, dto *CreateBookDTO) *book.Book
}
