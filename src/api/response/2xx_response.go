package response

import (
	"encoding/json"
	"net/http"
)

func SuccessfullyOk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]bool{
		"Message": true,
	})
}
