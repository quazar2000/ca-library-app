package composites

import (
	"github.com/quazar2000/ca-library-app/internal/adapters/api"

	"github.com/quazar2000/ca-library-app/internal/domain/book"

	book3 "github.com/quazar2000/ca-library-app/internal/adapters/db/book"

	book2 "github.com/quazar2000/ca-library-app/internal/adapters/api/book"
)

type BookComposite struct {
	Storage book.Storage
	Service book2.Service
	Handler api.Handler
}

func NewBookComposite(mongoComposite *MongoDBComposite) (*BookComposite, error) {
	storage := book3.NewStorage(mongoComposite.db)
	service := book.NewService(storage)
	handler := book2.NewHandler(service)
	return &BookComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil

}
