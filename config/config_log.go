package config

import (
	"log"
	"os"
)

// Fungsi untuk menyiapkan log file
func SetupLogFile() *os.File {
	// Memastikan folder log ada
	err := os.MkdirAll("log", os.ModePerm)
	if err != nil {
		log.Fatal("Gagal membuat folder log:", err)
		return nil
	}

	// Membuka atau membuat file log
	logFile, err := os.OpenFile("log/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Gagal membuka log file:", err)
		return nil
	}

	// Menetapkan log output ke file
	log.SetOutput(logFile)
	return logFile
}
