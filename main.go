package main

import "fmt"

func tampilkanMenu() {
	fmt.Println("\n==== MENU PINJAMAN ONLINE ====")
	fmt.Println("[1] Ajukan Pinjaman")
	fmt.Println("[2] Lihat Data Pinjaman Saya")
	fmt.Println("[3] Pelunasan")
	fmt.Println("[4] Urutkan Data Peminjam")
	fmt.Println("[5] Profil Saya")
	fmt.Println("[0] Keluar")
	fmt.Print("Pilih menu: ")
}

func main() {
	var (
		pilihan     int
		statusMasuk bool
		userAktif   *Pengguna
	)
	seed()

	for {
		pilihan = menuAwal()

		switch pilihan {
		case 1:
			for {
				userAktif = masuk(&dbPengguna)
				statusMasuk = userAktif != nil

				if !statusMasuk {
					var coba string
					fmt.Print("Coba lagi? (y/n): ")
					fmt.Scan(&coba)
					if coba != "y" && coba != "Y" {
						break
					}
				} else {
					break
				}
			}

			if statusMasuk {
				menuUtama(userAktif)
			}

		case 2:
			buatAkunBaru(&dbPengguna)

		case 0:
			fmt.Println("\nTerima kasih telah menggunakan TraPinjaman Online!")
			return

		default:
			fmt.Println("❌ Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func menuUtama(userAktif *Pengguna) {
	var pilihan int

	for {
		tampilkanMenu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			ajukanPinjaman(userAktif, &dbDataPeminjam)
		case 2:
			lihatDataPinjamanSaya(userAktif, &dbDataPeminjam)
		case 3:
			pelunasan(userAktif, &dbDataPeminjam)
		case 4:
			urutkanDataPeminjam(&dbDataPeminjam)
		case 5:
			menuProfil(userAktif)
		case 0:
			fmt.Println("\nTerima kasih, sampai jumpa lagi di TraPinjaman Online!")
			return
		default:
			fmt.Println("❌ Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func menuProfil(user *Pengguna) {
	var pilihan int

	for {
		fmt.Println("\n==== MENU PROFIL ====")
		fmt.Println("[1] Lihat Profil")
		fmt.Println("[2] Ubah Password")
		fmt.Println("[0] Kembali")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanProfil(user, &dbDataPeminjam)
		case 2:
			ubahPassword(user)
		case 0:
			return
		default:
			fmt.Println("❌ Pilihan tidak valid.")
		}
	}
}

func urutkanDataPeminjam(db *[100]Pinjaman) {
	fmt.Println("\n==== URUTKAN DATA PEMINJAM ====")
	fmt.Println("[1] Urutkan berdasarkan Jumlah Pinjaman")
	fmt.Println("[2] Urutkan berdasarkan Tenor")
	fmt.Println("[3] Urutkan berdasarkan Bunga")
	fmt.Println("[4] Urutkan berdasarkan Status")
	fmt.Println("[0] Kembali")
	fmt.Print("Pilih: ")

	var pilihUrut int
	fmt.Scan(&pilihUrut)

	switch pilihUrut {
	case 1:
		insertionSortByJumlah(db)
		tampilkanSemuaPinjaman(db)
	case 2:
		insertionSortByTenor(db)
		tampilkanSemuaPinjaman(db)
	case 3:
		insertionSortByBunga(db)
		tampilkanSemuaPinjaman(db)
	case 4:
		insertionSortByStatus(db)
		tampilkanSemuaPinjaman(db)
	case 0:
		return
	default:
		fmt.Println("❌ Pilihan tidak valid.")
	}
}

func formatRupiah(nilai int) string {
	var result string
	strNilai := fmt.Sprintf("%d", nilai)

	for i := len(strNilai) - 1; i >= 0; i-- {
		result = string(strNilai[i]) + result
		if (len(strNilai)-i)%3 == 0 && i > 0 {
			result = "." + result
		}
	}

	return "Rp" + result
}
