package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fungsi untuk membaca properti dari file
func loadProperties(filePath string) (map[string]string, error) {
	properties := make(map[string]string)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka file properti: %w", err) // Perbaiki pesan error menjadi huruf kecil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			properties[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("gagal membaca file properti: %w", err) // Perbaiki pesan error menjadi huruf kecil
	}

	return properties, nil
}

// Fungsi untuk mendapatkan koneksi URL dari properti
func GetDBURL() (string, error) {
	properties, err := loadProperties("config/db.properties")
	if err != nil {
		return "", err
	}
	return properties["db.url"], nil
}

// Fungsi untuk mendapatkan username dari properti
func GetDBUsername() (string, error) {
	properties, err := loadProperties("config/db.properties")
	if err != nil {
		return "", err
	}
	return properties["db.username"], nil
}

// Fungsi untuk mendapatkan password dari properti
func GetDBPassword() (string, error) {
	properties, err := loadProperties("config/db.properties")
	if err != nil {
		return "", err
	}
	return properties["db.password"], nil
}
