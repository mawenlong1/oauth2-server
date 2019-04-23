package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteJSON ...
func WriteJSON(w http.ResponseWriter, v interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

// Error ..
func Error(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": err})
}

// UnauthorizedError ...
func UnauthorizedError(w http.ResponseWriter, err string) {
	w.Header().Set("WWW-Authenticate", fmt.Sprintf("Bearer realm=%s", "oauth2-server"))
	Error(w, err, http.StatusUnauthorized)
}
