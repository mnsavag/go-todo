package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type listHandlers struct{}

func (h *listHandlers) GetRouter() chi.Router {
	lists := chi.NewRouter()

	lists.Post("/", h.createList)
	lists.Get("/", h.getAllLists)
	lists.Get("/{id}", h.getListById)
	lists.Put("/{id}", h.updateList)
	lists.Delete("/{id}", h.deleteList)

	return lists
}

func (h *listHandlers) createList(w http.ResponseWriter, r *http.Request) {

}

func (h *listHandlers) getAllLists(w http.ResponseWriter, r *http.Request) {

}

func (h *listHandlers) getListById(w http.ResponseWriter, r *http.Request) {

}

func (h *listHandlers) updateList(w http.ResponseWriter, r *http.Request) {

}

func (h *listHandlers) deleteList(w http.ResponseWriter, r *http.Request) {

}
