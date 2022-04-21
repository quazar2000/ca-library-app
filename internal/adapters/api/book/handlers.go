package user

import (
	"net/http"

	"github.com/quazar2000/ca-library-app/internal/adapters/api"

	"github.com/julienschmidt/httprouter"
)

const (
	bookURL  = "/books/:book_id"
	booksURL = "/books"
)

type handler struct {
	bookService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{bookService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(booksURL, h.GetAllBooks)
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// books := h.bookService.GetAllBooks(context.Bachground(), 0, 0)
	w.Write([]byte("books"))
	w.Header(http.StatusOK)
}
