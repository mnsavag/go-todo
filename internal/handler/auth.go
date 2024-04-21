package handler

import (
	"goTodo/internal/lib/server"
	"goTodo/internal/model"

	"encoding/json"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name     string `json:"name" validate:"required"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
		return
	}

	err := server.RequestBodyValidate(req)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
		return
	}

	user := model.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}

	// Добавить проверку что user уже есть в базе

	id, err := h.Services.Authorization.CreateUser(user)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusInternalServerError, err.Error(), "")
		return
	}

	server.Respond(w, r, http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
		return
	}

	err := server.RequestBodyValidate(req)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
		return
	}

	token, err := h.Services.Authorization.GenerateToken(req.Username, req.Password)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusInternalServerError, err.Error(), "")
		return
	}

	server.Respond(w, r, http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
