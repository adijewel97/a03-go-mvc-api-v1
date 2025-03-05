package middleware

import (
	"net/http"
	"strings"
)

// Token statis untuk autentikasi
const validToken = "550e8400-e29b-41d4-a716-446655440000"

// Middleware untuk memverifikasi token
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Memeriksa header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"error": "Authorization header missing"}`, http.StatusUnauthorized)
			return
		}

		// Menghapus prefix "Bearer " dari token
		token := strings.TrimPrefix(authHeader, "Bearer ")
		token = strings.TrimSpace(token) // Hapus spasi ekstra

		// Memeriksa token
		if token != validToken {
			http.Error(w, `{"error": "Invalid or missing token"}`, http.StatusUnauthorized)
			return
		}

		// Jika token valid, lanjutkan ke handler berikutnya
		next(w, r)
	}
}
