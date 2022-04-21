package author

import (
	"github.com/quazar2000/ca-library-app/internal/adapters/api"

	"github.com/julienschmidt/httprouter"
)

type handler struct {
	authorService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{authorService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	panic("implement me")
}
