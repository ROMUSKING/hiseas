package shared

import (
	"encoding/json"
	"net/http"
)

// JSONResponse standardizes API responses
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
