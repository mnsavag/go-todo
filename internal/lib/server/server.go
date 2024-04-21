package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func HttpErrResponse(w http.ResponseWriter, r *http.Request, code int, msg string, userMsg string) {
	Respond(w, r, code, map[string]interface{}{
		"message":     msg,
		"userMessage": userMsg,
	})
}

func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func RequestBodyValidate(requestBody interface{}) error {
	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		return err
	}
	return nil
}
