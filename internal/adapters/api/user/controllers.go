package user

import (
	adapters "github.com/quazar2000/ca-library-app/internal/adapters/user"
	"github.com/quazar2000/ca-library-app/internal/book"
)

type handler struct {
	bookService book.Service
}

func NewHandler(service book.Service) adapters.Handler {
	return &handler{bookService: service}
}
