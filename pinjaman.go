package main

import "fmt"

func ajukanPinjaman(userAktif *Pengguna, db *[100]Pinjaman) {
	var (
		nNominal       int
		nTenor         int
		nAngsuran      int
		nBunga         float64
		indeksTersedia int
	)

	fmt.Println("\n==== AJUKAN PINJAMAN ====")
	fmt.Print("Masukkan nominal pinjaman: ")
	fmt.Scan(&nNominal)

	if nNominal <= 0 {
		fmt.Println("❌ Nominal pinjaman harus lebih dari 0.")
		return
	}

	if nNominal > 100000000 {
		fmt.Println("❌ Maaf, maksimal pinjaman adalah Rp100.000.000.")
		return
	}

	pilihanTenor(nNominal)

	fmt.Print("Pilih tenor (dalam bulan/hari): ")
	fmt.Scan(&nTenor)
	fmt.Print("Pilih jumlah angsuran: ")
	fmt.Scan(&nAngsuran)

	if !validasiTenorAngsuran(nNominal, nTenor, nAngsuran) {
		fmt.Println("❌ Tenor atau jumlah angsuran tidak valid untuk nominal pinjaman ini.")
		return
	}

	nBunga = hitungBunga(nNominal)
	indeksTersedia = indeksKosongDbPeminjam(db)

	if indeksTersedia == -1 {
		fmt.Println("❌ Maaf, database pinjaman penuh.")
		return
	}

	dbDataPeminjam[indeksTersedia] = Pinjaman{
		idPeminjam:     userAktif.idPengguna,
		jumlahPinjaman: nNominal,
		tenor:          nTenor,
		bunga:          nBunga,
		jumlahAngsuran: nAngsuran,
		angsuranBayar:  0,
		statusLunas:    false,
	}

	fmt.Println("\n✅ Pengajuan pinjaman berhasil!")
	fmt.Printf("Jumlah Pinjaman: %s\n", formatRupiah(nNominal))
	fmt.Printf("Tenor: %d\n", nTenor)
	fmt.Printf("Bunga: %.2f%%\n", nBunga*100)
	fmt.Printf("Jumlah Angsuran: %d kali\n", nAngsuran)

	totalBunga := int(float64(nNominal) * nBunga)
	totalBayar := nNominal + totalBunga
	perAngsuran := totalBayar / nAngsuran

	fmt.Printf("Total Bunga: %s\n", formatRupiah(totalBunga))
	fmt.Printf("Total yang Harus Dibayar: %s\n", formatRupiah(totalBayar))
	fmt.Printf("Cicilan per Angsuran: %s\n", formatRupiah(perAngsuran))
}

func hitungBunga(nominal int) float64 {
	switch {
	case nominal < 10000000:
		return 0.05 
	case nominal <= 50000000:
		return 0.08 
	case nominal <= 100000000:
		return 0.15 
	default:
		return 0.18 
	}
}

func pilihanTenor(nominal int) {
	fmt.Printf("\nNominal Pinjaman: %s\n", formatRupiah(nominal))
	fmt.Println("Pilihan tenor dan cicilan yang tersedia:")

	switch {
	case nominal <= 300000:
		fmt.Println("- Tenor: 7 Hari, 14 Hari, 30 Hari")
		fmt.Println("- Cicilan: 1x, 2x, 3x")
	case nominal <= 1000000:
		fmt.Println("- Tenor: 30 Hari, 60 Hari, 90 Hari")
		fmt.Println("- Cicilan: 1x, 2x, 3x")
	case nominal <= 3000000:
		fmt.Println("- Tenor: 3 Bulan, 4 Bulan, 6 Bulan")
		fmt.Println("- Cicilan: 3x, 4x, 6x")
	case nominal <= 5000000:
		fmt.Println("- Tenor: 3 Bulan, 6 Bulan, 9 Bulan")
		fmt.Println("- Cicilan: 3x, 6x, 9x")
	case nominal <= 10000000:
		fmt.Println("- Tenor: 6 Bulan, 9 Bulan, 12 Bulan")
		fmt.Println("- Cicilan: 6x, 9x, 12x")
	case nominal <= 20000000:
		fmt.Println("- Tenor: 12 Bulan, 18 Bulan, 24 Bulan")
		fmt.Println("- Cicilan: 12x, 18x, 24x")
	case nominal <= 50000000:
		fmt.Println("- Tenor: 24 Bulan, 30 Bulan, 36 Bulan")
		fmt.Println("- Cicilan: 24x, 30x, 36x")
	default:
		fmt.Println("- Tenor: 36 Bulan, 48 Bulan, 60 Bulan")
		fmt.Println("- Cicilan: 36x, 48x, 60x")
	}
}

func validasiTenorAngsuran(nominal, tenor, angsuran int) bool {
	switch {
	case nominal <= 300000:
		return (tenor == 7 || tenor == 14 || tenor == 30) && (angsuran >= 1 && angsuran <= 3)
	case nominal <= 1000000:
		return (tenor == 30 || tenor == 60 || tenor == 90) && (angsuran >= 1 && angsuran <= 3)
	case nominal <= 3000000:
		return (tenor == 3 || tenor == 4 || tenor == 6) && (angsuran >= 3 && angsuran <= 6)
	case nominal <= 5000000:
		return (tenor == 3 || tenor == 6 || tenor == 9) && (angsuran >= 3 && angsuran <= 9)
	case nominal <= 10000000:
		return (tenor == 6 || tenor == 9 || tenor == 12) && (angsuran >= 6 && angsuran <= 12)
	case nominal <= 20000000:
		return (tenor == 12 || tenor == 18 || tenor == 24) && (angsuran >= 12 && angsuran <= 24)
	case nominal <= 50000000:
		return (tenor == 24 || tenor == 30 || tenor == 36) && (angsuran >= 24 && angsuran <= 36)
	default:
		return (tenor == 36 || tenor == 48 || tenor == 60) && (angsuran >= 36 && angsuran <= 60)
	}
}
