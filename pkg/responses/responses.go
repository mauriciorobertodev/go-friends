package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func Json(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

func Error(w http.ResponseWriter, status int, err error) {
	Json(w, status, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
