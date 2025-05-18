package main

import "fmt"

type Investasi struct {
	Nama          string
	Jenis         string
	Dana          float64
	NilaiSekarang float64
}

var investasiList []Investasi

func main() {
	for {
		var pilihan int
		fmt.Println("\n=== APLIKASI MANAJEMEN INVESTASI ===")
		fmt.Println("1. Menambahkan investasi")
		fmt.Println("2. Mengubah investasi")
		fmt.Println("3. Menghapus investasi")
		fmt.Println("4. Transaksi")
		fmt.Println("5. Portofolio")
		fmt.Println("6. Laporan / Analisis")
		fmt.Println("7. Keuntungan/Rugi")
		fmt.Println("8. Cari Investasi")
		fmt.Println("9. Keluar")
		fmt.Print("Pilih menu: ")

		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahInvestasi()
		case 2:
			ubahInvestasi()
		case 3:
			hapusInvestasi()
		case 4:
			updateNilai()
		case 5:
			tampilkanPortofolio()
		case 6:
			tampilkanLaporan()
		case 7:
			tampilkanKeuntungan()
		case 8:
			cariInvestasi()
		case 9:
			fmt.Println("Terima kasih telah menggunakan aplikasi.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahInvestasi() {
	var nama, jenis string
	var dana float64

	fmt.Print("Nama Investasi: ")
	fmt.Scanln(&nama)
	fmt.Print("Jenis (Saham/Crypto/ReksaDana): ")
	fmt.Scanln(&jenis)
	fmt.Print("Jumlah Dana: ")
	fmt.Scanln(&dana)

	investasiList = append(investasiList, Investasi{Nama: nama, Jenis: jenis, Dana: dana, NilaiSekarang: dana})
	fmt.Println("Investasi berhasil ditambahkan.")
}

func ubahInvestasi() {
	var nama string
	fmt.Print("Nama investasi yang ingin diubah: ")
	fmt.Scanln(&nama)

	for i := 0; i < len(investasiList); i++ {
		if investasiList[i].Nama == nama {
			var namaBaru, jenisBaru string
			var danaBaru float64

			fmt.Print("Nama baru: ")
			fmt.Scanln(&namaBaru)
			fmt.Print("Jenis baru: ")
			fmt.Scanln(&jenisBaru)
			fmt.Print("Dana baru: ")
			fmt.Scanln(&danaBaru)

			investasiList[i] = Investasi{
				Nama:          namaBaru,
				Jenis:         jenisBaru,
				Dana:          danaBaru,
				NilaiSekarang: danaBaru,
			}
			fmt.Println("Data berhasil diubah.")
			return
		}
	}
	fmt.Println("Investasi tidak ditemukan.")
}

func hapusInvestasi() {
	var nama string
	fmt.Print("Nama investasi yang ingin dihapus: ")
	fmt.Scanln(&nama)

	for i := 0; i < len(investasiList); i++ {
		if investasiList[i].Nama == nama {
			investasiList = append(investasiList[:i], investasiList[i+1:]...)
			fmt.Println("Investasi berhasil dihapus.")
			return
		}
	}
	fmt.Println("Investasi tidak ditemukan.")
}

func updateNilai() {
	var nama string
	var nilaiBaru float64
	fmt.Print("Nama investasi: ")
	fmt.Scanln(&nama)

	for i := 0; i < len(investasiList); i++ {
		if investasiList[i].Nama == nama {
			fmt.Print("Masukkan nilai saat ini: ")
			fmt.Scanln(&nilaiBaru)
			investasiList[i].NilaiSekarang = nilaiBaru
			fmt.Println("Nilai berhasil diperbarui.")
			return
		}
	}
	fmt.Println("Investasi tidak ditemukan.")
}

func tampilkanPortofolio() {
	fmt.Println("\n--- PORTOFOLIO ---")
	for i := 0; i < len(investasiList); i++ {
		fmt.Printf("Nama: %s | Jenis: %s | Dana: %.2f | Nilai Sekarang: %.2f\n", investasiList[i].Nama, investasiList[i].Jenis, investasiList[i].Dana, investasiList[i].NilaiSekarang)
	}
}

func tampilkanLaporan() {
	totalDana := 0.0
	totalNilai := 0.0

	for i := 0; i < len(investasiList); i++ {
		totalDana += investasiList[i].Dana
		totalNilai += investasiList[i].NilaiSekarang
	}

	keuntungan := totalNilai - totalDana
	persen := (keuntungan / totalDana) * 100

	fmt.Println("\n--- LAPORAN PORTOFOLIO ---")
	fmt.Printf("Total Dana: %.2f\n", totalDana)
	fmt.Printf("Nilai Sekarang: %.2f\n", totalNilai)
	fmt.Printf("Keuntungan/Rugi: %.2f (%.2f%%)\n", keuntungan, persen)
}

func tampilkanKeuntungan() {
	totalDana, totalNilai := 0.0, 0.0
	for i := 0; i < len(investasiList); i++ {
		totalDana += investasiList[i].Dana
		totalNilai += investasiList[i].NilaiSekarang
	}

	if validasikeuntungan(totalDana, totalNilai) {
		fmt.Println("Status: Menguntungkan")
	} else {
		fmt.Println("Status: Rugi")
	}
}

func validasikeuntungan(totalDana, totalNilai float64) bool {
	keuntungan := totalNilai - totalDana
	return keuntungan >= 0
}

func cariInvestasi() {
	var target string
	fmt.Print("Masukkan nama investasi yang dicari: ")
	fmt.Scanln(&target)

	for i := 0; i < len(investasiList); i++ {
		if investasiList[i].Nama == target {
			fmt.Printf("Ditemukan: %+v\n", investasiList[i])
			return
		}
	}
	fmt.Println("Tidak ditemukan.")
}
