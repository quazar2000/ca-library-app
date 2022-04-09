package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	adapters "github.com/quazar2000/ca-library-app/internal/adapters/api"
)

const (
	bookURL  = "/books/:book_id"
	booksURL = "/books"
)

type handler struct {
	bookService book.Service
}

func NewHandler(service book.Service) adapters.Handler {
	return &handler{bookService: service}
}

func (h *handler) REgister(router *httprouter.Router) {
	router.GET(booksURL, h.GetAllBooks)
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// books := h.bookService.GetAllBooks(context.Bachground(), 0, 0)
	w.Write([]byte("books"))
	w.Header(http.StatusOK)
}
