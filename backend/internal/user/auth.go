package user

import (
	"net/http"
	"os"
	"strings"
)

// Simple token-based authentication middleware for zero trust
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if !validateToken(token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Token validation logic (replace with real implementation)
func validateToken(token string) bool {
	// Example: Bearer <token>
	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return false
	}
	// Compare with a secret token from env (for demo)
	return parts[1] == os.Getenv("API_TOKEN")
}
