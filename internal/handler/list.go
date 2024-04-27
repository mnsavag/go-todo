package handler

import (
	"encoding/json"
	"goTodo/internal/lib/server"
	"goTodo/internal/model"

	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) createList(w http.ResponseWriter, r *http.Request) {
	type dto struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
	}
	reqBody := &dto{}
	if err := server.ParseReqToDto(w, r, reqBody); err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
		return
	}

	todo := model.TodoList{
		Title:       reqBody.Title,
		Description: reqBody.Description,
	}

	userId, err := getUserId(r)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
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
	userId, err := getUserId(r)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
		return
	}

	lists, err := h.Services.TodoList.GetAll(userId)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusInternalServerError, err.Error(), "")
		return
	}

	server.Respond(w, r, http.StatusOK, map[string]interface{}{
		"data": lists,
	})
}

func (h *Handler) getListById(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
		return
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, "invalid id param", "")
		return
	}

	list, err := h.Services.TodoList.GetById(userId, id)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusInternalServerError, err.Error(), "")
		return
	}

	if list == (model.TodoList{}) {
		server.Respond(w, r, http.StatusOK, map[string]interface{}{})
		return
	}
	server.Respond(w, r, http.StatusOK, list)
}

func (h *Handler) updateList(w http.ResponseWriter, r *http.Request) {
	var todoInput model.UpdateListInput

	if err := json.NewDecoder(r.Body).Decode(&todoInput); err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
	}

	userId, err := getUserId(r)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
		return
	}
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, "invalid id param", "")
		return
	}

	err = h.Services.TodoList.Update(userId, id, todoInput)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
		return
	}

	server.Respond(w, r, http.StatusOK, map[string]interface{}{})
}

func (h *Handler) deleteList(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
		return
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, "invalid id param", "")
		return
	}

	err = h.Services.TodoList.Delete(userId, id)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusInternalServerError, err.Error(), "")
		return
	}

	server.Respond(w, r, http.StatusOK, map[string]interface{}{})
}
