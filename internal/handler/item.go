package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type itemHandlers struct{}

func (h *itemHandlers) GetRouter() chi.Router {
	items := chi.NewRouter()

	items.Post("/", h.createItem)
	items.Get("/", h.getAllItems)
	items.Get("/{item_id}", h.getItemById)
	items.Put("/{item_id}", h.updateItem)
	items.Delete("/{item_id}", h.deleteItem)

	return items
}

func (h *itemHandlers) createItem(w http.ResponseWriter, r *http.Request) {

}

func (h *itemHandlers) getAllItems(w http.ResponseWriter, r *http.Request) {

}

func (h *itemHandlers) getItemById(w http.ResponseWriter, r *http.Request) {

}

func (h *itemHandlers) updateItem(w http.ResponseWriter, r *http.Request) {

}

func (h *itemHandlers) deleteItem(w http.ResponseWriter, r *http.Request) {

}
