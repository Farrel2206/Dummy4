package main

func seed() {
	dbPengguna[0] = Pengguna{1, "Agha", "agha@example.com", "password123"}
	dbPengguna[1] = Pengguna{2, "Elfan", "elfangamtenk@example.com", "elfan123"}
	dbPengguna[2] = Pengguna{3, "Citra", "citra@example.com", "citra456"}
	dbPengguna[3] = Pengguna{4, "Farel", "farel67@example.com", "farel456"}
	dbPengguna[4] = Pengguna{5, "Geby", "geby22@example.com", "geby456"}
	dbPengguna[5] = Pengguna{6, "Caca", "cacaboom@example.com", "caca456"}
	dbPengguna[6] = Pengguna{7, "Geebry", "geebry@example.com", "1234567"}
	dbPengguna[7] = Pengguna{8, "Admin", "admin@trapinjaman.com", "admin123"}
	dbPengguna[8] = Pengguna{9, "Demo", "demo@trapinjaman.com", "demo123"}

	dbDataPeminjam[0] = Pinjaman{1, 1000000, 12, 0.05, 12, 0, false}
	dbDataPeminjam[1] = Pinjaman{2, 2000000, 24, 0.08, 24, 24, true}
	dbDataPeminjam[2] = Pinjaman{3, 1500000, 18, 0.06, 18, 6, false}
	dbDataPeminjam[3] = Pinjaman{4, 2000000, 18, 0.06, 18, 0, false}
	dbDataPeminjam[4] = Pinjaman{5, 2500000, 18, 0.06, 18, 3, false}
	dbDataPeminjam[5] = Pinjaman{6, 1000000, 12, 0.05, 12, 12, true}
	dbDataPeminjam[6] = Pinjaman{7, 1200000, 12, 0.05, 12, 4, false}
	dbDataPeminjam[7] = Pinjaman{1, 500000, 6, 0.05, 6, 2, false}
	dbDataPeminjam[8] = Pinjaman{3, 3000000, 24, 0.08, 24, 12, false}
	dbDataPeminjam[9] = Pinjaman{8, 2000000, 12, 0.06, 12, 0, false}
}
