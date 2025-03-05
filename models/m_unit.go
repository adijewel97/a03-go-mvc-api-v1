package models

import (
	"a03-my-go-project/config"
	"database/sql"
	"log"
)

type UnitData struct {
	KD_DIST     string `json:"kd_dist"`
	NAMA_DIST   string `json:"nama_dist"`
	KD_AREA     string `json:"kd_area"`
	AREA        string `json:"area"`
	NAMA_AREA   string `json:"nama_area"`
	KD_UNIT     string `json:"kd_unit"`
	UNITUP      string `json:"unitup"`
	NAMA        string `json:"nama"`
	NAMA_UNIT   string `json:"nama_unit"`
	ALAMAT      string `json:"alamat"`
	ALAMAT_UNIT string `json:"alamat_unit"`
}

type UnitFilter struct {
	KD_DIST string `json:"kd_dist,omitempty"`
	KD_AREA string `json:"kd_area,omitempty"`
	KD_UNIT string `json:"kd_unit,omitempty"`
}

func GetMstUnitData(db *sql.DB, filter UnitFilter) ([]UnitData, error) {

	// Menyiapkan log file
	logFile := config.SetupLogFile()
	defer logFile.Close() // Pastikan log file ditutup setelah selesai

	query := `SELECT KD_DIST, NAMA_DIST, KD_AREA, AREA, NAMA_AREA, KD_UNIT, UNITUP, NAMA, 
			  NAMA_UNIT, ALAMAT, ALAMAT_UNIT FROM MST_UNIT WHERE 1=1`
	var args []interface{}

	// Tambahkan filter sesuai permintaan
	if filter.KD_DIST != "" {
		query += " AND KD_DIST = :1"
		args = append(args, filter.KD_DIST)
	}
	if filter.KD_AREA != "" {
		query += " AND KD_AREA = :2"
		args = append(args, filter.KD_AREA)
	}
	if filter.KD_UNIT != "" {
		query += " AND KD_UNIT = :3"
		args = append(args, filter.KD_UNIT)
	}

	// Menambahkan log untuk query yang dijalankan
	log.Println("Query yang dijalankan:", query)

	rows, err := db.Query(query, args...)
	if err != nil {
		// Menambahkan log jika terjadi error
		log.Println("Error menjalankan query:", err)
		return nil, err
	}
	defer rows.Close()

	var units []UnitData
	for rows.Next() {
		var unit UnitData
		if err := rows.Scan(&unit.KD_DIST, &unit.NAMA_DIST, &unit.KD_AREA, &unit.AREA,
			&unit.NAMA_AREA, &unit.KD_UNIT, &unit.UNITUP, &unit.NAMA, &unit.NAMA_UNIT,
			&unit.ALAMAT, &unit.ALAMAT_UNIT); err != nil {
			log.Println("Error saat scan row:", err)
			return nil, err
		}
		units = append(units, unit)
	}

	// Menambahkan log setelah berhasil mendapatkan data
	log.Println("Data berhasil diambil:", len(units), "unit")

	return units, nil
}
