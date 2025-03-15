package save

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	random "restapi/internal/randomfunc"

	"github.com/go-playground/validator"
)

type URLSaver interface {
	SaveURL(urlForSave string, name string) error
}
type Request struct {
	URL  string `json:"url" validate:"required,url"`
	Name string `json:"name,omitempty"`
}
type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Name   string `json:"name,omitempty"`
}

const lenght = 6

func New(log *slog.Logger, urlSave URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if errors.Is(err, io.EOF) {
			log.Info("Request empty")
			json.NewEncoder(w).Encode(&Response{
				Status: "StatusError",
				Error:  "request is empty",
			})
			return
		}
		if err != nil {
			log.Info("Decode tofailed")
			json.NewEncoder(w).Encode(&Response{
				Status: "StatusError",
				Error:  "failed to decode",
			})
			return
		}

		if err := validator.New().Struct(req); err != nil {
			validateErr := err.(validator.ValidationErrors)
			log.Info("validated to failed")
			json.NewEncoder(w).Encode(&Response{
				Status: "StatusError",
				Error:  validateErr.Error(),
			})
			return
		}
		name := req.Name
		if name == "" {
			name = random.RandomStrings(lenght)
		}
		err = urlSave.SaveURL(req.URL, name)
		if err != nil {
			log.Info("DB save error")
			json.NewEncoder(w).Encode(&Response{
				Status: "StatusError",
				Error:  "Failed to saveUrl in DB",
			})
			return
		}
		answerOk(w, r, name)
	}
}
func answerOk(w http.ResponseWriter, r *http.Request, name string) {
	json.NewEncoder(w).Encode(&Response{
		Status: "StatusOK",
		Name:   name,
	})
}
