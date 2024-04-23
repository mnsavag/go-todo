package handler

import (
	"goTodo/internal/lib/server"
	"goTodo/internal/model"
	"net/http"
)

func (h *Handler) createList(w http.ResponseWriter, r *http.Request) {
	type dto struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
	reqBody := &dto{}
	if err := server.RequestValidate(w, r, reqBody); err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
		return
	}

	todo := model.TodoList{
		Title:       reqBody.Title,
		Description: reqBody.Description,
	}

	userId, err := getUserId(r)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, "user id not found", "")
		return
	}

	listId, err := h.Services.TodoList.Create(userId, todo)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusInternalServerError, err.Error(), "")
		return
	}

	server.Respond(w, r, http.StatusOK, map[string]interface{}{
		"id": listId,
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
