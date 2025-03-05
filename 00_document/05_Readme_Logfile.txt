Untuk menambahkan log file di folder log, Anda dapat menggunakan package log yang disediakan oleh Go, dan mengonfigurasi log agar mencatat informasi ke file di direktori tertentu.

Berikut adalah langkah-langkah untuk menambahkan pencatatan log ke file di folder log dalam aplikasi Anda.

Menambahkan Log ke File: Anda bisa mengonfigurasi package log untuk menulis log ke file dengan menggunakan os.OpenFile untuk membuka (atau membuat) file log.

Menambahkan Log di dalam Fungsi GetMstUnitData: Di dalam fungsi ini, kita bisa mencatat query yang dijalankan dan juga mencatat error atau informasi lain ke file log.

Berikut adalah contoh modifikasi kode Anda dengan logging ke file:

Langkah 1: Menambahkan Logging ke File
Tambahkan fungsi untuk menyiapkan log file:

go
Copy
Edit
package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// Fungsi untuk menyiapkan log file
func setupLogFile() *os.File {
	// Membuka atau membuat file log
	logFile, err := os.OpenFile("log/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Gagal membuka log file:", err)
		return nil
	}
	// Menetapkan log output ke file
	log.SetOutput(logFile)
	return logFile
}

func GetMstUnitData(db *sql.DB, filter UnitFilter) ([]UnitData, error) {
	// Memastikan file log siap
	logFile := setupLogFile()
	defer logFile.Close() // Pastikan file ditutup saat selesai

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
Penjelasan:
setupLogFile: Fungsi ini membuka (atau membuat jika tidak ada) file log di folder log dengan nama app.log. File ini akan mencatat log secara append (menambah data ke file yang ada).

Logging: Di dalam fungsi GetMstUnitData, setiap kali ada query yang dijalankan atau jika ada error, log akan dicatat ke file tersebut. Misalnya:

Setelah query dijalankan, kita log query-nya.
Jika terjadi error saat query atau saat memproses row, log akan mencatatnya.
Ketika data berhasil diambil, log juga mencatat jumlah unit yang berhasil diambil.
Langkah 2: Struktur Folder
Pastikan Anda memiliki folder log di direktori kerja proyek Anda. Jika belum ada, buatlah folder tersebut agar log dapat ditulis ke file app.log.

Langkah 3: Verifikasi Log
Setelah Anda menjalankan aplikasi dan melakukan request, Anda akan menemukan file log di folder log dengan informasi tentang query yang dijalankan, error (jika ada), dan informasi sukses.

Dengan begitu, Anda dapat melacak aktivitas API Anda melalui file log di folder log. aa