package main

import "fmt"

func lihatDataPinjamanSaya(user *Pengguna, db *[100]Pinjaman) {
	indeksPinjaman := cariPinjamanById(db, user.idPengguna)

	if len(indeksPinjaman) == 0 {
		fmt.Println("\n❗ Anda belum memiliki pinjaman.")
		return
	}

	fmt.Printf("\n==== DATA PINJAMAN %s ====\n", user.nama)
	fmt.Printf("Jumlah pinjaman aktif: %d\n", hitungJumlahPinjamanAktif(db, user.idPengguna))
	fmt.Printf("Total pinjaman aktif: %s\n", formatRupiah(hitungTotalPinjaman(db, user.idPengguna)))
	fmt.Println("\n=== DAFTAR PINJAMAN ===")

	for i, idx := range indeksPinjaman {
		pinjaman := db[idx]
		statusStr := "Belum Lunas"
		if pinjaman.statusLunas {
			statusStr = "Lunas"
		}

		fmt.Printf("\n[%d] Pinjaman #%d\n", i+1, idx+1)
		fmt.Printf("   Jumlah Pinjaman: %s\n", formatRupiah(pinjaman.jumlahPinjaman))
		fmt.Printf("   Tenor: %d\n", pinjaman.tenor)
		fmt.Printf("   Bunga: %.2f%%\n", pinjaman.bunga*100)
		fmt.Printf("   Angsuran: %d kali\n", pinjaman.jumlahAngsuran)
		fmt.Printf("   Angsuran Terbayar: %d kali\n", pinjaman.angsuranBayar)
		fmt.Printf("   Status: %s\n", statusStr)

		if !pinjaman.statusLunas {
			totalBunga := hitungTotalBungaPinjaman(pinjaman)
			totalBayar := pinjaman.jumlahPinjaman + totalBunga
			perAngsuran := hitungJumlahAngsuran(pinjaman)
			sisaAngsuran := pinjaman.jumlahAngsuran - pinjaman.angsuranBayar

			fmt.Printf("   Total Bunga: %s\n", formatRupiah(totalBunga))
			fmt.Printf("   Total Harus Dibayar: %s\n", formatRupiah(totalBayar))
			fmt.Printf("   Cicilan per Angsuran: %s\n", formatRupiah(perAngsuran))
			fmt.Printf("   Sisa Angsuran: %d kali (%s)\n",
				sisaAngsuran, formatRupiah(perAngsuran*sisaAngsuran))
		}

		fmt.Println("   ---------------------------")
	}
}

func pelunasan(user *Pengguna, db *[100]Pinjaman) {
	indeksPinjaman := cariPinjamanById(db, user.idPengguna)
	pinjamanAktif := false

	for _, idx := range indeksPinjaman {
		if !db[idx].statusLunas {
			pinjamanAktif = true
			break
		}
	}

	if !pinjamanAktif {
		fmt.Println("\n❗ Anda tidak memiliki pinjaman aktif yang perlu dilunasi.")
		return
	}

	fmt.Println("\n==== MENU PELUNASAN ====")
	fmt.Println("[1] Bayar Angsuran")
	fmt.Println("[2] Pelunasan Langsung")
	fmt.Println("[0] Kembali")
	fmt.Print("Pilih: ")

	var pilihan int
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		bayarAngsuran(user, db)
	case 2:
		pelunasanLangsung(user, db)
	case 0:
		return
	default:
		fmt.Println("❌ Pilihan tidak valid.")
	}
}

func bayarAngsuran(user *Pengguna, db *[100]Pinjaman) {
	var pinjamanAktif []int
	var nomorPinjaman int

	fmt.Println("\n==== BAYAR ANGSURAN ====")
	fmt.Println("Pinjaman aktif Anda:")

	// Menampilkan pinjaman aktif
	count := 1
	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam == user.idPengguna && !db[i].statusLunas && db[i].idPeminjam != 0 {
			perAngsuran := hitungJumlahAngsuran(db[i])

			fmt.Printf("[%d] Pinjaman #%d - %s (Cicilan: %s per angsuran)\n",
				count, i+1, formatRupiah(db[i].jumlahPinjaman), formatRupiah(perAngsuran))

			pinjamanAktif = append(pinjamanAktif, i)
			count++
		}
	}

	if len(pinjamanAktif) == 0 {
		fmt.Println("❗ Anda tidak memiliki pinjaman aktif.")
		return
	}

	fmt.Print("Pilih pinjaman (0 untuk kembali): ")
	fmt.Scan(&nomorPinjaman)

	if nomorPinjaman == 0 || nomorPinjaman > len(pinjamanAktif) {
		return
	}

	idx := pinjamanAktif[nomorPinjaman-1]
	perAngsuran := hitungJumlahAngsuran(db[idx])

	fmt.Printf("\nDetail Pinjaman #%d\n", idx+1)
	fmt.Printf("Jumlah Pinjaman: %s\n", formatRupiah(db[idx].jumlahPinjaman))
	fmt.Printf("Cicilan per Angsuran: %s\n", formatRupiah(perAngsuran))
	fmt.Printf("Angsuran Terbayar: %d dari %d\n", db[idx].angsuranBayar, db[idx].jumlahAngsuran)

	var jumlahBayar int
	fmt.Print("\nMasukkan jumlah angsuran yang ingin dibayar: ")
	fmt.Scan(&jumlahBayar)

	if jumlahBayar <= 0 {
		fmt.Println("❌ Jumlah angsuran harus lebih dari 0.")
		return
	}

	sisaAngsuran := db[idx].jumlahAngsuran - db[idx].angsuranBayar

	if jumlahBayar > sisaAngsuran {
		fmt.Printf("❌ Jumlah angsuran maksimal adalah %d.\n", sisaAngsuran)
		return
	}

	// Proses pembayaran
	db[idx].angsuranBayar += jumlahBayar
	totalBayar := perAngsuran * jumlahBayar

	// Cek jika sudah lunas
	if db[idx].angsuranBayar >= db[idx].jumlahAngsuran {
		db[idx].statusLunas = true
		fmt.Println("\n🎉 Selamat! Pinjaman telah LUNAS!")
	} else {
		sisaAngsuran = db[idx].jumlahAngsuran - db[idx].angsuranBayar
		fmt.Printf("\n✅ Pembayaran berhasil! Sisa %d angsuran lagi.\n", sisaAngsuran)
	}

	fmt.Printf("Total Pembayaran: %s\n", formatRupiah(totalBayar))
}

func pelunasanLangsung(user *Pengguna, db *[100]Pinjaman) {
	var pinjamanAktif []int
	var nomorPinjaman int

	fmt.Println("\n==== PELUNASAN LANGSUNG ====")
	fmt.Println("Pinjaman aktif Anda:")

	// Menampilkan pinjaman aktif
	count := 1
	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam == user.idPengguna && !db[i].statusLunas && db[i].idPeminjam != 0 {
			perAngsuran := hitungJumlahAngsuran(db[i])
			sisaAngsuran := db[i].jumlahAngsuran - db[i].angsuranBayar
			totalSisa := perAngsuran * sisaAngsuran

			fmt.Printf("[%d] Pinjaman #%d - %s (Sisa: %s)\n",
				count, i+1, formatRupiah(db[i].jumlahPinjaman), formatRupiah(totalSisa))

			pinjamanAktif = append(pinjamanAktif, i)
			count++
		}
	}

	if len(pinjamanAktif) == 0 {
		fmt.Println("❗ Anda tidak memiliki pinjaman aktif.")
		return
	}

	fmt.Print("Pilih pinjaman (0 untuk kembali): ")
	fmt.Scan(&nomorPinjaman)

	if nomorPinjaman == 0 || nomorPinjaman > len(pinjamanAktif) {
		return
	}

	idx := pinjamanAktif[nomorPinjaman-1]
	perAngsuran := hitungJumlahAngsuran(db[idx])
	sisaAngsuran := db[idx].jumlahAngsuran - db[idx].angsuranBayar
	totalSisa := perAngsuran * sisaAngsuran

	fmt.Printf("\nDetail Pinjaman #%d\n", idx+1)
	fmt.Printf("Jumlah Pinjaman: %s\n", formatRupiah(db[idx].jumlahPinjaman))
	fmt.Printf("Sisa Angsuran: %d dari %d\n", sisaAngsuran, db[idx].jumlahAngsuran)
	fmt.Printf("Total yang Harus Dibayar: %s\n", formatRupiah(totalSisa))

	var konfirmasi string
	fmt.Print("\nApakah Anda yakin ingin melunasi pinjaman ini? (y/n): ")
	fmt.Scan(&konfirmasi)

	if konfirmasi == "y" || konfirmasi == "Y" {
		// Proses pelunasan
		db[idx].angsuranBayar = db[idx].jumlahAngsuran
		db[idx].statusLunas = true

		fmt.Println("\n🎉 Selamat! Pinjaman telah LUNAS!")
		fmt.Printf("Total Pembayaran: %s\n", formatRupiah(totalSisa))
	} else {
		fmt.Println("❌ Pelunasan dibatalkan.")
	}
}
