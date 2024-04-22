package handler

import (
	"goTodo/internal/lib/server"
	"net/http"
)

func (h *Handler) createList(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(userCtx)
	server.Respond(w, r, http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllLists(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getListById(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateList(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteList(w http.ResponseWriter, r *http.Request) {

}
