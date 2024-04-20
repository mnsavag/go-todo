package handler

import (
	"goTodo/internal/lib/server"
	"goTodo/internal/model"

	"encoding/json"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	req := &request{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		server.HttpErr(w, r, http.StatusBadRequest, err)
		return
	}

	user := model.User{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}

	id, err := h.Services.Authorization.CreateUser(user)
	if err != nil {
		server.HttpErr(w, r, http.StatusInternalServerError, err)
		return
	}

	server.Respond(w, r, http.StatusCreated, id)

	// валидация
	// унифицировать ответ сервера в успешном случае
	// убрать константу op, так как она только для разработки
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {

}
