package handler

import (
	"goTodo/internal/lib/server"

	"context"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			server.HttpErrResponse(w, r, http.StatusUnauthorized, "empty auth header", "user is not authorized")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			server.HttpErrResponse(w, r, http.StatusUnauthorized, "invalid auth header", "user is not authorized")
			return
		}

		userId, err := h.Services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			server.HttpErrResponse(w, r, http.StatusUnauthorized, err.Error(), "")
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
