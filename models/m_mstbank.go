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
	All        bool   `json:"all,omitempty"` // Tambahkan field All untuk menampilkan semua data
}

// Fungsi untuk mengambil data dari MST_BANK dengan filter
func GetMstBankData(db *sql.DB, filter BankFilter) ([]BankData, error) {
	query := "SELECT CDC_B_ID, CDC_B_NAME, CDC_B_REGISTERED FROM MST_BANK WHERE 1=1"
	var args []interface{}

	// Cek jika filter "all" diaktifkan
	if filter.All {
		// Tidak ada filter, ambil semua data
		query = "SELECT CDC_B_ID, CDC_B_NAME, CDC_B_REGISTERED FROM MST_BANK"
	} else {
		// Menambahkan kondisi filter untuk CDC_B_ID jika ada
		if filter.CDC_B_ID != "" {
			query += " AND CDC_B_ID = ?"
			args = append(args, filter.CDC_B_ID)
		}
		// Menambahkan kondisi filter untuk CDC_B_NAME jika ada
		if filter.CDC_B_NAME != "" {
			query += " AND UPPER(CDC_B_NAME) = UPPER(?)"
			args = append(args, filter.CDC_B_NAME)
		}
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

	// Mengembalikan hasil data
	return banks, nil
}
