Bagus, kamu sudah berhasil memecah API ke dalam bentuk model, controller, dan pengaturan database! Berikut adalah saran struktur file yang bisa kamu gunakan untuk memecah API tersebut:

curl -X POST http://localhost:8080/mstunit/select -d '{"kd_dist": "53"}' -H "Content-Type: application/json"



Struktur Direktori:
bash
Copy
Edit
/project-root
    /controllers
        mstbank.go
    /models
        mstbank.go
    /db
        db.go
    main.go
1. main.go (Masih tetap menjadi entry point)
go
Copy
Edit
package main

import (
	"log"
	"net/http"
	"your_project/controllers"
)

func main() {
	http.HandleFunc("/mstbank/select", controllers.SelectMstBankData)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
2. /controllers/mstbank.go (Controller untuk menangani logika API)
go
Copy
Edit
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"your_project/models"
	"your_project/db"
)

// Fungsi untuk mengambil data dari MST_BANK
func SelectMstBankData(w http.ResponseWriter, r *http.Request) {
	var filter models.BankFilter
	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		http.Error(w, "Gagal membaca data JSON", http.StatusBadRequest)
		return
	}

	// Koneksi ke database
	database, err := db.Connect()
	if err != nil {
		http.Error(w, "Gagal koneksi ke database", http.StatusInternalServerError)
		return
	}
	defer database.Close()

	// Menjalankan query dengan filter
	banks, err := models.GetMstBankData(database, filter)
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal mengambil data: %v", err), http.StatusInternalServerError)
		return
	}

	// Menyusun response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(banks)
}
3. /models/mstbank.go (Model untuk menangani data dan query)
go
Copy
Edit
package models

import "database/sql"

// Struktur untuk menampung data tabel MST_BANK
type BankData struct {
	CDC_B_ID         string `json:"cdc_b_id"`
	CDC_B_NAME       string `json:"cdc_b_name"`
	CDC_B_REGISTERED string `json:"cdc_b_registered"`
}

// Struktur untuk request parameter filter
type BankFilter struct {
	CDC_B_ID   string `json:"cdc_b_id,omitempty"`
	CDC_B_NAME string `json:"cdc_b_name,omitempty"`
}

// Fungsi untuk mengambil data dari MST_BANK dengan filter
func GetMstBankData(db *sql.DB, filter BankFilter) ([]BankData, error) {
	query := "SELECT CDC_B_ID, CDC_B_NAME, CDC_B_REGISTERED FROM MST_BANK WHERE 1=1"
	var args []interface{}

	// Menambahkan kondisi filter
	if filter.CDC_B_ID != "" {
		query += " AND CDC_B_ID = :1"
		args = append(args, filter.CDC_B_ID)
	}
	if filter.CDC_B_NAME != "" {
		query += " AND CDC_B_NAME = :2"
		args = append(args, filter.CDC_B_NAME)
	}

	// Menjalankan query
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Menyimpan hasil query ke dalam slice
	var banks []BankData
	for rows.Next() {
		var bank BankData
		if err := rows.Scan(&bank.CDC_B_ID, &bank.CDC_B_NAME, &bank.CDC_B_REGISTERED); err != nil {
			return nil, err
		}
		banks = append(banks, bank)
	}

	return banks, nil
}
4. /db/db.go (Untuk pengaturan koneksi database)
go
Copy
Edit
package db

import (
	"database/sql"
	_ "github.com/godror/godror"
	"fmt"
)

// Fungsi untuk menghubungkan ke database
func Connect() (*sql.DB, error) {
	dsn := "USERADISPPOBNTL/adis123@127.0.0.1:1521/adis.iconpln.co.id"
	db, err := sql.Open("godror", dsn)
	if err != nil {
		return nil, fmt.Errorf("Gagal koneksi ke database: %w", err)
	}
	return db, nil
}
Dengan memecah kode seperti ini, kamu akan lebih mudah mengelola proyek, terutama ketika aplikasi berkembang. Kamu bisa menambah fungsionalitas atau membuat unit-test untuk setiap bagian dengan lebih mudah.