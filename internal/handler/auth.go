package handler

import (
	"goTodo/internal/lib/server"
	"goTodo/internal/model"

	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	type dto struct {
		Name     string `json:"name" validate:"required"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	reqBody := &dto{}
	if err := server.ParseReqToDto(w, r, reqBody); err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
		return
	}

	user := model.User{
		Name:     reqBody.Name,
		Username: reqBody.Username,
		Password: reqBody.Password,
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
	type dto struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	reqBody := &dto{}
	if err := server.ParseReqToDto(w, r, reqBody); err != nil {
		server.HttpErrResponse(w, r, http.StatusBadRequest, err.Error(), "")
		return
	}

	token, err := h.Services.Authorization.GenerateToken(reqBody.Username, reqBody.Password)
	if err != nil {
		server.HttpErrResponse(w, r, http.StatusInternalServerError, err.Error(), "")
		return
	}

	server.Respond(w, r, http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
