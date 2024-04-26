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

func ParseReqToDto(w http.ResponseWriter, r *http.Request, dto interface{}) error {
	/* decode json into dto and validate body according to dto*/
	if err := json.NewDecoder(r.Body).Decode(dto); err != nil {
		return err
	}
	err := RequestBodyValidate(dto)
	if err != nil {
		return err
	}
	return nil
}

func RequestBodyValidate(requestBody interface{}) error {
	validate := validator.New()
	if err := validate.Struct(requestBody); err != nil {
		return err
	}
	return nil
}
