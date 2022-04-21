package composites

import (
	"github.com/quazar2000/ca-library-app/internal/adapters/api"

	author2 "github.com/quazar2000/ca-library-app/internal/domain/author"

	author3 "github.com/quazar2000/ca-library-app/internal/adapters/api/author"

	"github.com/quazar2000/ca-library-app/internal/adapters/api/author"
)

type AuthorComposite struct {
	Storage author2.Storage
	Service author.Service
	Handler api.Handler
}

func NewAuthorComposite(mongoComposite *MongoDBComposite) (*AuthorComposite, error) {
	storage := author3.NewStorage(mongoComposite.db)
	service := author2.NewService(storage)
	handler := author.NewHandler(service)
	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil

}
