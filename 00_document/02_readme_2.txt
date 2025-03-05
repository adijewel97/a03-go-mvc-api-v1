ada aplikasi yang harus diinstall untuk golang connectString ke Oracle

 1. Persyaratan Saat Build (Build-time Requirements)
    Agar bisa build proyek yang menggunakan Godror, dibutuhkan:

    Go 1.15 atau lebih baru
    Pastikan kamu menggunakan Go 1.21.x (karena Godror mungkin tidak kompatibel dengan Go 1.24).
    C Compiler dengan CGO_ENABLED=1
    Karena Godror menggunakan C bindings ke Oracle Client, kamu harus mengaktifkan CGO dan punya C compiler (gcc) yang terinstal.
    Cross-compilation (build untuk platform lain) sulit dilakukan karena bergantung pada sistem C.

    C:\TDM-GCC-64\bin

    module a03-my-go-project
    versi diturnkan
      go 1.21.13

      require github.com/godror/godror v0.47.0

      require (
         github.com/go-logfmt/logfmt v0.6.0 // indirect
         github.com/godror/knownpb v0.1.2 // indirect
         golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 // indirect
         google.golang.org/protobuf v1.34.2 // indirect
      )

