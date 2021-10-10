package response

import (
	"encoding/json"
	"net/http"
)

func ServerError(w http.ResponseWriter, r *http.Request, res map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	json.NewEncoder(w).Encode(res)
}
