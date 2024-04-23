package handler

import (
	"errors"
	"net/http"
)

func getUserId(r *http.Request) (int64, error) {
	id := r.Context().Value(userCtx)
	if id == nil {
		return 0, errors.New("user id not found")
	}

	idInt64, ok := id.(int64)
	if !ok {
		return 0, errors.New("user id not found")
	}

	return idInt64, nil
}
