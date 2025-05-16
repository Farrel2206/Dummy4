package main

import "fmt"

func insertionSortByJumlah(db *[100]Pinjaman) {
	for i := 1; i < len(db); i++ {
		key := db[i]
		j := i - 1
		if db[i].idPeminjam == 0 {
			continue
		}
		for j >= 0 && db[j].jumlahPinjaman > key.jumlahPinjaman {
			db[j+1] = db[j]
			j--
		}
		db[j+1] = key
	}
}

func insertionSortByTenor(db *[100]Pinjaman) {
	for i := 1; i < len(db); i++ {
		key := db[i]
		j := i - 1
		if db[i].idPeminjam == 0 {
			continue
		}
		for j >= 0 && db[j].tenor > key.tenor {
			db[j+1] = db[j]
			j--
		}
		db[j+1] = key
	}
}

func tampilkanSemuaPinjaman(db *[100]Pinjaman) {
	fmt.Println("\n==== DAFTAR PINJAMAN TERURUT ====")

	count := 0
	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam != 0 {
			count++
			statusStr := "Belum Lunas"
			if db[i].statusLunas {
				statusStr = "Lunas"
			}

			fmt.Printf("\n[%d] Pinjaman #%d\n", count, i+1)
			fmt.Printf("   ID Peminjam: %d\n", db[i].idPeminjam)
			fmt.Printf("   Jumlah Pinjaman: %s\n", formatRupiah(db[i].jumlahPinjaman))
			fmt.Printf("   Tenor: %d\n", db[i].tenor)
			fmt.Printf("   Bunga: %.2f%%\n", db[i].bunga*100)
			fmt.Printf("   Angsuran: %d dari %d\n", db[i].angsuranBayar, db[i].jumlahAngsuran)
			fmt.Printf("   Status: %s\n", statusStr)
			fmt.Println("   ---------------------------")
		}
	}

	if count == 0 {
		fmt.Println("❗ Belum ada data pinjaman.")
	}
}

func insertionSortByBunga(db *[100]Pinjaman) {
	for i := 1; i < len(db); i++ {
		key := db[i]
		j := i - 1
		if db[i].idPeminjam == 0 {
			continue
		}
		for j >= 0 && db[j].bunga > key.bunga {
			db[j+1] = db[j]
			j--
		}
		db[j+1] = key
	}
}

func insertionSortByStatus(db *[100]Pinjaman) {
	for i := 1; i < len(db); i++ {
		key := db[i]
		j := i - 1
		if db[i].idPeminjam == 0 {
			continue
		}
		for j >= 0 && (db[j].statusLunas && !key.statusLunas) {
			db[j+1] = db[j]
			j--
		}
		db[j+1] = key
	}
}
