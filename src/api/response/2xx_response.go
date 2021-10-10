package response

import (
	"encoding/json"
	"net/http"
)

func SuccessfullyOk(w http.ResponseWriter, r *http.Request, res map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(res)
}

func SuccessfullyCreated(w http.ResponseWriter, r *http.Request, res map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(res)
}
