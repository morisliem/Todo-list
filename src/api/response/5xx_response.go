package response

import (
	"net/http"
)

func ServerError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
}
