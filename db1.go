package main

type Pinjaman struct {
	idPeminjam     int
	jumlahPinjaman int
	tenor          int
	bunga          float64
	jumlahAngsuran int
	angsuranBayar  int 
	statusLunas    bool
}

type Pengguna struct {
	idPengguna int
	nama       string
	email      string
	password   string
}

var dbDataPeminjam [100]Pinjaman
var dbPengguna [100]Pengguna

func indeksKosongDbPeminjam(db *[100]Pinjaman) int {
	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam == 0 {
			return i
		}
	}
	return -1
}

func indeksKosongDbPengguna(db *[100]Pengguna) int {
	for i := 0; i < len(db); i++ {
		if db[i].idPengguna == 0 {
			return i
		}
	}
	return -1
}

func cariPinjamanById(db *[100]Pinjaman, idPeminjam int) []int {
	var hasil []int
	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam == idPeminjam && db[i].idPeminjam != 0 {
			hasil = append(hasil, i)
		}
	}
	return hasil
}

func hitungTotalPinjaman(db *[100]Pinjaman, idPeminjam int) int {
	total := 0
	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam == idPeminjam && !db[i].statusLunas {
			total += db[i].jumlahPinjaman
		}
	}
	return total
}

func hitungJumlahPinjamanAktif(db *[100]Pinjaman, idPeminjam int) int {
	jumlah := 0
	for i := 0; i < len(db); i++ {
		if db[i].idPeminjam == idPeminjam && !db[i].statusLunas {
			jumlah++
		}
	}
	return jumlah
}

func hitungTotalBungaPinjaman(pinjaman Pinjaman) int {
	return int(float64(pinjaman.jumlahPinjaman) * pinjaman.bunga)
}

func hitungJumlahAngsuran(pinjaman Pinjaman) int {
	totalBunga := hitungTotalBungaPinjaman(pinjaman)
	totalBayar := pinjaman.jumlahPinjaman + totalBunga
	return totalBayar / pinjaman.jumlahAngsuran
}
