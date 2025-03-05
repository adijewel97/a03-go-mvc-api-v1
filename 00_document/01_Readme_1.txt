Berikut langkah-langkah untuk menginstal dan membuat contoh koneksi ke Oracle menggunakan Godror di Golang:

1. Pastikan Oracle Instant Client Sudah Terpasang
    Sudah diinstal di:
    makefile

    D:\app\adi.setiadi\product\12.2.0
    Pastikan path sudah ditambahkan ke environment variable Windows:
    makefile

    D:\app\adi.setiadi\product\12.2.0\bin
    Bisa diuji dengan menjalankan di Command Prompt:
    bash

    sqlplus USERADISPPOBNTL/adis123@//localhost:1521/adis.iconpln.co.id
    Jika bisa masuk ke SQLPlus, berarti Instant Client sudah siap digunakan.
2. Instal Golang dan Godror
    Pastikan Go sudah terinstal dengan mengecek versi:
    go

    go version
    Buat folder proyek, lalu jalankan:
    swift

    go mod init example.com/godror_test
    go get github.com/godror/godror

    -- go get github.com/godror/godror@latest
3. Buat File main.go
    Buat file main.go dengan isi sebagai berikut:

    go

    package main

    import (
        "context"
        "database/sql"
        "fmt"
        "log"

        _ "github.com/godror/godror"
    )

    func main() {
        // Konfigurasi koneksi sesuai dengan environment Laravel yang sudah ada
        dsn := `user="USERADISPPOBNTL" password="adis123" connectString="//localhost:1521/adis.iconpln.co.id"`

        db, err := sql.Open("godror", dsn)
        if err != nil {
            log.Fatalf("Gagal membuka koneksi: %v", err)
        }
        defer db.Close()

        // Uji koneksi
        ctx := context.Background()
        err = db.PingContext(ctx)
        if err != nil {
            log.Fatalf("Gagal terhubung ke database: %v", err)
        }

        fmt.Println("Koneksi berhasil!")

        // Contoh query sederhana
        var sysDate string
        err = db.QueryRowContext(ctx, "SELECT TO_CHAR(SYSDATE, 'YYYY-MM-DD HH24:MI:SS') FROM dual").Scan(&sysDate)
        if err != nil {
            log.Fatalf("Gagal menjalankan query: %v", err)
        }

        fmt.Println("Tanggal & Waktu DB:", sysDate)
    }
4. Jalankan Program
    Jalankan perintah berikut untuk menjalankan program:

    go

    go run main.go
    Jika berhasil, outputnya akan seperti ini:

    yaml

    Koneksi berhasil!
    Tanggal & Waktu DB: 2025-03-04 12:34:56
    Silakan coba dan beri tahu jika ada kendala! 🚀




