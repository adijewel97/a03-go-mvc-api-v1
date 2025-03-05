package controllers

import (
	"a03-my-go-project/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// Fungsi untuk mengambil data dari MST_BANK
// Fungsi untuk mengambil data dari MST_BANK
func SelectMstBankData(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Memeriksa apakah metode HTTP adalah POST
		if r.Method != http.MethodPost {
			http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
			return
		}

		var filter models.BankFilter
		err := json.NewDecoder(r.Body).Decode(&filter)
		if err != nil {
			http.Error(w, "Gagal membaca data JSON", http.StatusBadRequest)
			return
		}

		// Menjalankan query dengan filter menggunakan models.GetMstBankData
		banks, err := models.GetMstBankData(db, filter)
		if err != nil {
			http.Error(w, fmt.Sprintf("Gagal mengambil data: %v", err), http.StatusInternalServerError)
			return
		}

		// Menyusun response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(banks)
	}
}
