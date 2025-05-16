package main

import (
	"fmt"
	"strings"
)

func menuAwal() int {
	var pilihan int

	fmt.Println("\n====== TRAPINJAMAN ONLINE ======")
	fmt.Println("======= SELAMAT DATANG! =======")
	fmt.Println("[1] Login")
	fmt.Println("[2] Buat Akun Baru")
	fmt.Println("[0] Keluar")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)

	return pilihan
}

func buatAkunBaru(db *[100]Pengguna) bool {
	var (
		nNama     string
		nEmail    string
		nPassword string
		idBaru    int = 1
	)

	fmt.Println("\n====== BUAT AKUN BARU ======")
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nNama)
	fmt.Print("Masukkan email: ")
	fmt.Scan(&nEmail)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&nPassword)

	// Validasi email unik
	for i := 0; i < len(db); i++ {
		if db[i].email == nEmail && db[i].idPengguna != 0 {
			fmt.Println("\n❌ Email sudah digunakan. Silakan gunakan email lain.")
			return false
		}
		// Cari ID terbesar untuk menentukan ID baru
		if db[i].idPengguna >= idBaru {
			idBaru = db[i].idPengguna + 1
		}
	}

	// Cari indeks kosong untuk pengguna baru
	indeksKosong := -1
	for i := 0; i < len(db); i++ {
		if db[i].idPengguna == 0 {
			indeksKosong = i
			break
		}
	}

	if indeksKosong == -1 {
		fmt.Println("\n❌ Database pengguna penuh. Tidak bisa membuat akun baru.")
		return false
	}

	// Simpan pengguna baru
	db[indeksKosong] = Pengguna{
		idPengguna: idBaru,
		nama:       nNama,
		email:      nEmail,
		password:   nPassword,
	}

	fmt.Println("\n✅ Akun berhasil dibuat! Silakan login dengan akun baru Anda.")
	return true
}

func masuk(db *[100]Pengguna) *Pengguna {
	var (
		nNama     string
		nPassword string
	)

	fmt.Println("\n====== LOGIN ======")
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nNama)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&nPassword)

	for i := 0; i < len(db); i++ {
		if strings.EqualFold(nNama, db[i].nama) && nPassword == db[i].password {
			fmt.Println("\n✅ Login berhasil!")
			fmt.Printf("Selamat datang, %s!\n", db[i].nama)
			return &db[i]
		}
	}

	fmt.Println("\n❌ Login gagal. Username atau password salah.")
	return nil
}

func tampilkanProfil(user *Pengguna, db *[100]Pinjaman) {
	fmt.Println("\n====== PROFIL SAYA ======")
	fmt.Println("=======================")
	fmt.Printf("ID Pengguna : %d\n", user.idPengguna)
	fmt.Printf("Nama        : %s\n", user.nama)
	fmt.Printf("Email       : %s\n", user.email)
	fmt.Println("=======================")

	// Menampilkan data pinjaman
	fmt.Println("\n===== DATA PINJAMAN =====")

	totalPinjaman := hitungTotalPinjaman(db, user.idPengguna)
	pinjamanAktif := hitungJumlahPinjamanAktif(db, user.idPengguna)

	fmt.Printf("Total Pinjaman Aktif: %s\n", formatRupiah(totalPinjaman))
	fmt.Printf("Jumlah Pinjaman Aktif: %d\n", pinjamanAktif)

	fmt.Println("\nRincian Pinjaman:")

	indeksPinjaman := cariPinjamanById(db, user.idPengguna)

	if len(indeksPinjaman) == 0 {
		fmt.Println("Anda belum memiliki pinjaman.")
	} else {
		for i, idx := range indeksPinjaman {
			pinjaman := db[idx]
			statusStr := "Belum Lunas"
			if pinjaman.statusLunas {
				statusStr = "Lunas"
			}

			fmt.Printf("\n[%d] Pinjaman #%d\n", i+1, idx+1)
			fmt.Printf("   Jumlah: %s\n", formatRupiah(pinjaman.jumlahPinjaman))

			if !pinjaman.statusLunas {
				sisaAngsuran := pinjaman.jumlahAngsuran - pinjaman.angsuranBayar
				perAngsuran := hitungJumlahAngsuran(pinjaman)
				totalSisa := perAngsuran * sisaAngsuran

				fmt.Printf("   Sisa Angsuran: %d kali (%s)\n",
					sisaAngsuran, formatRupiah(totalSisa))
			}

			fmt.Printf("   Status: %s\n", statusStr)
		}
	}
	fmt.Println("=======================")
}

func ubahPassword(user *Pengguna) {
	var (
		passwordLama string
		passwordBaru string
	)

	fmt.Print("\nMasukkan password lama: ")
	fmt.Scan(&passwordLama)

	if passwordLama != user.password {
		fmt.Println("❌ Password lama tidak sesuai!")
		return
	}

	fmt.Print("Masukkan password baru: ")
	fmt.Scan(&passwordBaru)

	user.password = passwordBaru
	fmt.Println("✅ Password berhasil diubah!")
}
