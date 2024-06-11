package handler

import (
	"goTodo/internal/lib/server"
	"goTodo/internal/model"

	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)

	listId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
	}

	type dto struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
	reqBody := &dto{}
	if err := server.ParseReqToDto(w, r, reqBody); err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
		return
	}

	item := model.TodoItem{
		Title:       reqBody.Title,
		Description: reqBody.Description,
	}

	itemId, err := h.Services.TodoItem.Create(userId, listId, item)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusInternalServerError, err.Error(), "")
		return
	}

	server.Respond(w, r, http.StatusOK, map[string]interface{}{
		"id": itemId,
	})
}

func (h *Handler) getAllItems(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getItemById(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateItem(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteItem(w http.ResponseWriter, r *http.Request) {

}
