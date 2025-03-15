package redirect

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
)

type URLGetter interface {
	GetURL(name string) (string, error)
}
type Request struct {
	Name string `json:"name"`
}
type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	URL    string `json:"url,omitempty"`
}

func New(log *slog.Logger, urlgetter URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if name == "" {
			log.Info("name is empty")
			json.NewEncoder(w).Encode(&Response{
				Status: "StatusError",
				Error:  "faild with encoding dates",
				URL:    "",
			})
		}
		url, err := urlgetter.GetURL(name)
		if err != nil {
			log.Info("get url to failed")
			json.NewEncoder(w).Encode(&Response{
				Status: "StatusError",
				Error:  "DB to failed",
				URL:    "",
			})
		}
		http.Redirect(w, r, url, http.StatusFound)
	}
}
