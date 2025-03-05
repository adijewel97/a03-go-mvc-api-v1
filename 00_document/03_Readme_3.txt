go mod init a03-my-go-projrct


install dan buat contoh godror di golang

anbil contoh envorment laravel yang sudah jalan
#Menggunkan ke db langsung tanpa tns client
DB_CONNECTION=oracle
DB_HOST=localhost
DB_PORT=1521
DB_SERVICE_NAME=adis.iconpln.co.id
DB_USERNAME=USERADISPPOBNTL
DB_PASSWORD=adis123

instatall instan client yang sesuai dengan oracle 12 c 2 adalah
D:\app\adi.setiadi\product\12.2.0
path env windows path = D:\app\adi.setiadi\product\12.2.0\bin

dsn := "USERADISPPOBNTL/adis123@(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=localhost)(PORT=1521))(CONNECT_DATA=(SERVER=dedicated)(SERVICE_NAME=adis.iconpln.co.id)))"


ADIS=
  (DESCRIPTION=
    (ADDRESS=
      (PROTOCOL=TCP)
      (HOST=localhost)
      (PORT=1521)
    )
    (CONNECT_DATA=
      (SERVER=dedicated)
      (SERVICE_NAME=adis.iconpln.co.id)
    )
  )
  
  
  package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/godror/godror"
)

// Struktur untuk menampung data tabel
type TableData struct {
	TableName string `json:"table_name"`
}

func main() {
	http.HandleFunc("/tables", getTables)
	fmt.Println("Server berjalan di port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getTables(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("godror", "USERADISPPOBNTL/adis123@127.0.0.1:1521/adis.iconpln.co.id")
	if err != nil {
		http.Error(w, "Gagal koneksi ke database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Perbaikan query
	rows, err := db.Query("SELECT table_name FROM user_tables")
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal mengambil data: %v", err), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tables []TableData
	for rows.Next() {
		var table TableData
		if err := rows.Scan(&table.TableName); err != nil {
			http.Error(w, "Gagal membaca data", http.StatusInternalServerError)
			return
		}
		tables = append(tables, table)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tables)
}
sudah berhasil buatkan API untuk menggambil data select * from MST_BANK