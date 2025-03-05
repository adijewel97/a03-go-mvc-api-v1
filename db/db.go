package db

import (
	"a03-my-go-project/config" // Pastikan untuk mengimpor package config
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

// Fungsi untuk menghubungkan ke database
func Connect() (*sql.DB, error) {
	// Ambil konfigurasi dari file db.properties yang ada di folder /config
	dbURL, err := config.GetDBURL()
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan URL database: %w", err) // Perbaiki pesan error menjadi huruf kecil
	}

	dbUsername, err := config.GetDBUsername()
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan username database: %w", err) // Perbaiki pesan error menjadi huruf kecil
	}

	dbPassword, err := config.GetDBPassword()
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan password database: %w", err) // Perbaiki pesan error menjadi huruf kecil
	}

	// Menyusun DSN
	dsn := fmt.Sprintf("%s/%s@%s", dbUsername, dbPassword, dbURL)

	// Membuka koneksi ke database
	db, err := sql.Open("godror", dsn)
	if err != nil {
		return nil, fmt.Errorf("gagal koneksi ke database: %w", err) // Perbaiki pesan error menjadi huruf kecil
	}
	return db, nil
}
