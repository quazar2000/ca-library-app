package author

import (
	"context"

	"github.com/quazar2000/ca-library-app/internal/domain/author"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *author.Author
	GetAll(ctx context.Context, limit, offset int) []*author.Author
	Create(ctx context.Context, dto *CreateAuthorDTO) *author.Author
}
