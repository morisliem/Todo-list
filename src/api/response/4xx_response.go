package response

import (
	"encoding/json"
	"net/http"
)

func BadRequest(w http.ResponseWriter, r *http.Request, res map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	json.NewEncoder(w).Encode(res)
}

func NotFound(w http.ResponseWriter, r *http.Request, res map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(res)
}
