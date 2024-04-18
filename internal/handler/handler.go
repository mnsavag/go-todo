package handler

import (
	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Mount("/api/lists", new(listHandlers).GetRouter())
	router.Mount("/api/lists/{id}/items", new(itemHandlers).GetRouter())

	return router
}
