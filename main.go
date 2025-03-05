package main

import (
	"fmt"
	"log"
	"net/http"

	"a03-my-go-project/controllers"
	"a03-my-go-project/db"
	"a03-my-go-project/middleware"

	_ "github.com/godror/godror"
)

func main() {
	// Koneksi ke database
	db, err := db.Connect()
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}
	defer db.Close()

	// Tambahkan route dengan middleware autentikasi
	http.HandleFunc("/mstbank/select", middleware.Authenticate(controllers.SelectMstBankData(db)))
	http.HandleFunc("/mstunit/select", middleware.Authenticate(controllers.GetMstUnitHandler(db)))

	fmt.Println("Server berjalan di port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
