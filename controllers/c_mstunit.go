package controllers

import (
	"a03-my-go-project/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Fungsi untuk validasi token sederhana
func isValidToken(token string) bool {
	const validToken = "550e8400-e29b-41d4-a716-446655440000" // Token harus cocok
	return strings.TrimSpace(token) == validToken
}

func GetMstUnitHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Cek header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"error": "Authorization header is required"}`, http.StatusUnauthorized)
			return
		}

		// Format token harus "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, `{"error": "Invalid Authorization format"}`, http.StatusUnauthorized)
			return
		}

		token := strings.TrimSpace(tokenParts[1])
		fmt.Println("Token yang diterima:", token) // Debugging

		// Validasi token
		if !isValidToken(token) {
			http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
			return
		}

		// Decode JSON dari request body
		var filter models.UnitFilter
		err := json.NewDecoder(r.Body).Decode(&filter)
		if err != nil {
			http.Error(w, `{"error": "Gagal membaca data JSON"}`, http.StatusBadRequest)
			return
		}

		// Menjalankan query
		units, err := models.GetMstUnitData(db, filter)
		if err != nil {
			http.Error(w, fmt.Sprintf(`{"error": "Gagal mengambil data: %v"}`, err), http.StatusInternalServerError)
			return
		}

		// Menyusun response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(units)
	}
}
