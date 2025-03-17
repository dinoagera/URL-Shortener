package delhand

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

type DeleteURL interface {
	DeleteURL(name string) error
}
type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func New(log *slog.Logger, deleteurl DeleteURL) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			log.Info("url not found")
			json.NewEncoder(w).Encode(&Response{
				Status: "StatusError",
				Error:  "URL not found",
			})
		}
		err := deleteurl.DeleteURL(name)
		if err != nil {
			log.Info("db to failed")
			json.NewEncoder(w).Encode(&Response{
				Status: "StatusError",
				Error:  "URL not found",
			})
		}
		json.NewEncoder(w).Encode(&Response{
			Status: "StatusOK",
		})
	}
}
