//Tubes alpro 2
//klo
package main

import "fmt"

type project struct {
	nama      string
	kategori  string
	teknologi [10]string 
	countTech int        
	day       int
	bulan     int
	tahun     int
	kesulitan int
}

var daftarProject [100]project 
var countProject int           

func tambahProject() {
	if countProject >= 100 {
		fmt.Println("Database project penuh!")
		return
	}

	var p project
	fmt.Print("\nNama proyek: ")
	fmt.Scan(&p.nama)

	fmt.Print("Kategori: ")
	fmt.Scan(&p.kategori)

	fmt.Print("Masukkan jumlah teknologi digunakan (maks 10): ")
	fmt.Scan(&p.countTech)
	if p.countTech > 10 || p.countTech < 1 {
		fmt.Println()
		fmt.Println("Pilihan tidak valid!")
		return
	}

	fmt.Println("Teknologi yang digunakan: ")
	for i := 0; i < p.countTech; i++ {
		fmt.Printf("%d.  ", i+1)
		fmt.Scan(&p.teknologi[i])
	}

	fmt.Print("Masukkan tanggal (9 2 1991): ")
	fmt.Scan(&p.day, &p.bulan, &p.tahun)
	if p.day < 1 || p.day > 30 {
		fmt.Printf("\nData tidak tersimpan karena tanggal yang anda inputkan tidak valid\n")
		return
	}
	if p.bulan < 1 || p.bulan > 12 {
		fmt.Printf("\nData tidak tersimpan karena bulan yang anda inputkan tidak valid\n")
		return
	}

	fmt.Print("Tingkat kesulitan (1-5): ")
	fmt.Scan(&p.kesulitan)

	if p.kesulitan < 1 || p.kesulitan > 5 {
		fmt.Printf("\nData tidak tersimpan karena tingkat kesulitan yang anda masukkan diluar interval\n")
		return
	}

	daftarProject[countProject] = p
	countProject++
	fmt.Printf("\nProject berhasil ditambahkan!\n")
}

func editProject() {
	if countProject == 0 {
		fmt.Printf("\nBelum ada project yang tersimpan!\n")
		return
	}

	fmt.Printf("\nDaftar Project:\n")
	for i := 0; i < countProject; i++ {
		fmt.Printf("%d. %s\n", i+1, daftarProject[i].nama)
	}

	var nomor int
	fmt.Print("Pilih nomor project yang akan diedit: ")
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > countProject {
		fmt.Print("\nNomor project tidak valid!\n")
		return
	}
	var p *project
	var perbaikan string
	p = &daftarProject[nomor-1]

	fmt.Print("1. Nama\n2. Kategori\n3. Teknologi\n4. Tanggal\n5. Tingkat kesulitan\nBagian yang ingin diperbaiki: ")
	fmt.Scan(&perbaikan)

	if perbaikan == "1" {
		fmt.Print("Edit nama project menjadi: ", p.nama)
		fmt.Scan(&p.nama)
	} else if perbaikan == "2" {
		fmt.Print("Edit kategori project menjadi: ", p.kategori)
		fmt.Scan(&p.kategori)
	} else if perbaikan == "3" {
		var editanTech int
		var tambahanTech string
		var tujuanPerbaikan int
		var perbaikanTech string

		fmt.Println("\nDaftar teknologi sebelumnya: ")
		for i := 0; i < p.countTech; i++ {
			fmt.Printf("- %s\n", p.teknologi[i])
		}
		fmt.Print("1. Tambahkan teknologi\n2. Ganti teknologi\nPilihan: ")
		fmt.Scan(&editanTech)

		if editanTech == 1 {
			var jmlTambahanTech int
			var maxTambah int

			maxTambah = 10 - p.countTech
			if maxTambah <= 0 {
				fmt.Println("\nTeknologi sudah penuh, tidak bisa menambhakan teknologi baru")
				return
			}
			fmt.Printf("Berapa teknologi yang akan ditambahkan (%d): ", maxTambah)
			fmt.Scan(&jmlTambahanTech)

			if jmlTambahanTech > maxTambah {
				fmt.Println("Jumlah teknologi yang ingin anda tambahkan melebihi batas, Proses dibatalkan")
				return
			}

			for j := 0; j < jmlTambahanTech; j++ {
				fmt.Printf("%d. ", j+1)
				fmt.Scan(&tambahanTech)

				p.teknologi[p.countTech+j] = tambahanTech
			}
			p.countTech += jmlTambahanTech
		} else if editanTech == 2 {
			var i int
			fmt.Printf("Daftar teknologi yang digunakan: ")
			for i = 0; i < p.countTech; i++ {
				fmt.Printf("\n%d. %s", i+1, p.teknologi[i])
			}
			fmt.Printf("\nTeknologi mana yang ingin diperbaiki: ")
			fmt.Scan(&tujuanPerbaikan)

			if i < 1 || i > p.countTech {
				fmt.Println("Nomor tidak valid.")
				return
			}

			fmt.Printf("Ganti dengan: ")
			fmt.Scan(&perbaikanTech)
			p.teknologi[tujuanPerbaikan-1] = perbaikanTech
		} else {
			fmt.Println("\nPilihan tidak valid!")
			return
		}

	} else if perbaikan == "4" {
		fmt.Printf("\nEdit tanggal menjadi: ")
		fmt.Scan(&p.day, &p.bulan, &p.tahun)

	} else if perbaikan == "5" {
		fmt.Print("Edit tingkat kesulitan menjadi: ", p.kesulitan)
		fmt.Scan(&p.kesulitan)
	}
	fmt.Printf("\nProject berhasil diedit!\n")
}

func hapusProject() {
	if countProject == 0 {
		fmt.Println("Belum ada project yang tersimpan.")
		return
	}

	fmt.Println("Daftar Project:")
	for i := 0; i < countProject; i++ {
		fmt.Printf("%d. %s\n", i+1, daftarProject[i].nama)
	}

	var nomor int
	fmt.Print("Pilih nomor project yang akan dihapus: ")
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > countProject {
		fmt.Printf("\nNomor project tidak valid!\n")
		return
	}

	for i := nomor - 1; i < countProject-1; i++ {
		daftarProject[i] = daftarProject[i+1]
	}
	countProject--
	fmt.Println("Project berhasil dihapus!")
}

func keahlian() {
	if countProject == 0 {
		fmt.Println("Belum ada data teknologi.")
		return
	}

	var namaTeknologi [100]string
	var totalTeknologi int
	var techCount [100]int

	for i := 0; i < countProject; i++ {
		var sudahDihitung [100]string
		var countLokal int

		for j := 0; j < daftarProject[i].countTech; j++ {
			skip := false
			for k := 0; k < countLokal; k++ {
				if sudahDihitung[k] == daftarProject[i].teknologi[j] {
					skip = true
					break
				}
			}
			if skip {
				continue
			}
			sudahDihitung[countLokal] = daftarProject[i].teknologi[j]
			countLokal++

			found := false
			for k := 0; k < totalTeknologi; k++ {
				if namaTeknologi[k] == daftarProject[i].teknologi[j] {
					techCount[k]++
					found = true
					break
				}
			}
			if !found && totalTeknologi < 100 {
				namaTeknologi[totalTeknologi] = daftarProject[i].teknologi[j]
				techCount[totalTeknologi] = 1
				totalTeknologi++
			}
		}
	}
	fmt.Println("\nKeahlian yang telah dipelajari berdasarkan: ")
	for i := 0; i < totalTeknologi; i++ {
		fmt.Printf("%d. %s: %d project\n", i+1, namaTeknologi[i], techCount[i])
	}
	fmt.Println()
}

func cariProject() {
	var kataKunci string
	var pilihan, berdasarkan int

	fmt.Print("\nCari berdasarkan:\n1. Nama\n2. Kategori\nPilihan: ")
	fmt.Scan(&berdasarkan)
	fmt.Print("Masukkan kata kunci pencarian: ")
	fmt.Scan(&kataKunci)

	fmt.Print("\nCari menggunakan:\n1. Sequential Search\n2. Binary Search\nPilihan: ")
	fmt.Scan(&pilihan)

	var hasil [100]project
	var countHasil int

	if berdasarkan == 1 && pilihan == 1 {
		sequentialSearchNama(kataKunci, &hasil, &countHasil)
	} else if berdasarkan == 1 && pilihan == 2 {
		binarySearchByName(kataKunci, &hasil, &countHasil)
	} else if berdasarkan == 2 && pilihan == 1 {
		sequentialSearchKategori(kataKunci, &hasil, &countHasil)
	} else if berdasarkan == 2 && pilihan == 2 {
		fmt.Println("Binary search hanya tersedia untuk pencarian berdasarkan nama!")
		return
	} else {
		fmt.Println("Input tidak valid")
		return
	}

	tampilkanHasilPencarian(hasil, countHasil)
}

func sequentialSearchNama(kataKunci string, hasil *[100]project, countHasil *int) {
	*countHasil = 0
	for i := 0; i < countProject; i++ {
		if cekSubstring(daftarProject[i].nama, kataKunci) {
			hasil[*countHasil] = daftarProject[i]
			*countHasil += 1
		}
	}
}

func sequentialSearchKategori(kataKunci string, hasil *[100]project, countHasil *int) {
	*countHasil = 0
	for i := 0; i < countProject; i++ {
		if cekSubstring(daftarProject[i].kategori, kataKunci) {
			hasil[*countHasil] = daftarProject[i]
			*countHasil += 1
		}
	}
}

func binarySearchByName(kataKunci string, hasil *[100]project, countHasil *int) {
	mengurutkanNama()
	*countHasil = 0

	left := 0
	right := countProject - 1
	found := -1

	for left <= right {
		mid := (left + right) / 2
		if kataKunci == daftarProject[mid].nama {
			found = mid
			break
		} else if kataKunci < daftarProject[mid].nama {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if found == -1 {
		return
	}

	i := found
	for i >= 0 && daftarProject[i].nama == kataKunci {
		hasil[*countHasil] = daftarProject[i]
		*countHasil += 1
		i--
	}

	i = found + 1
	for i < countProject && daftarProject[i].nama == kataKunci {
		hasil[*countHasil] = daftarProject[i]
		*countHasil += 1
		i++
	}
}

func cekSubstring(teks, pattern string) bool {
	n := len(teks)
	m := len(pattern)

	if m == 0 || m > n {
		return false
	}

	for i := 0; i <= n-m; i++ {
		j := 0
		for j < m && teks[i+j] == pattern[j] {
			j++
		}
		if j == m {
			return true
		}
	}
	return false
}

func mengurutkanNama() {
	for i := 0; i < countProject-1; i++ {
		for j := 0; j < countProject-i-1; j++ {
			if daftarProject[j].nama > daftarProject[j+1].nama {
				temp := daftarProject[j]
				daftarProject[j] = daftarProject[j+1]
				daftarProject[j+1] = temp
			}
		}
	}
}

func tampilkanHasilPencarian(hasil [100]project, countHasil int) {
	if countHasil == 0 {
		fmt.Println("Tidak ditemukan project yang sesuai")
		return
	}

	fmt.Println("\nHasil Pencarian:")
	for i := 0; i < countHasil; i++ {
		p := hasil[i]
		fmt.Printf("%d. Nama: %s\n", i+1, p.nama)
		fmt.Println("   Kategori:", p.kategori)
		fmt.Print("   Teknologi: ")
		for j := 0; j < p.countTech; j++ {
			fmt.Print(p.teknologi[j])
			if j < p.countTech-1 {
				fmt.Print(", ")
			}
		}
		fmt.Printf("\n   Tanggal: %d/%d/%d\n", p.day, p.bulan, p.tahun)
		fmt.Println("   Kesulitan:", p.kesulitan)
		fmt.Println()
	}
}
func urutkanProject() {
	if countProject == 0 {
		fmt.Println("Belum ada project yang tersimpan.")
		return
	}

	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. Tingkat kesulitan")
	fmt.Println("2. Tanggal pembuatan")

	var pilihan int
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)
	if pilihan < 1 || pilihan > 2 {
		fmt.Println("Pilihan tidak valid!")
		return
	}
	var naikTurun int
	fmt.Print("1. Ascending (terkecil ke terbesar)\n2. Descending (terbesar ke terkecil) ")
	fmt.Print("\nPilihan: ")
	fmt.Scan(&naikTurun)

	if pilihan == 1 {
		if naikTurun == 1 {
			ascendingKesulitan()
		} else if naikTurun == 2 {
			descendingKesulitan()
		} else {
			fmt.Println("Pilihan tidak valid!")
			return
		}
	} else if pilihan == 2 {
		if naikTurun == 1 {
			ascendingTanggal()
		} else if naikTurun == 2 {
			descendingTanggal()
		} else {
			fmt.Println("Pilihan tidak valid!")
			return
		}
	}

	fmt.Println("Daftar project setelah diurutkan:")
	for i := 0; i < countProject; i++ {
		p := daftarProject[i]
		fmt.Println(i+1, ". Nama project: ", p.nama)
		fmt.Println("    Kategori: ", p.kategori)
		fmt.Print("    Teknologi yang digunakan: ")
		for i := 0; i < p.countTech; i++ {
			fmt.Print(p.teknologi[i])
			j := 1
			if j < p.countTech-1 {
				fmt.Print(", ")
				j++
			}
			fmt.Print(j)
		}
		fmt.Printf("\n    Tanggal pembuatan project: %d/%d/%d\n", p.day, p.bulan, p.tahun)
		fmt.Println("    Tingkat kesulitan: ", p.kesulitan)
	}
}
func ascendingKesulitan() {
	var pass, idx, i int
	var temp project

	pass = 1
	for pass <= countProject-1 {
		idx = pass - 1
		i = pass

		for i < countProject {
			if daftarProject[idx].kesulitan > daftarProject[i].kesulitan {
				idx = i
			}
			i = i + 1
		}
		temp = daftarProject[pass-1]
		daftarProject[pass-1] = daftarProject[idx]
		daftarProject[idx] = temp

		pass = pass + 1
	}

}
func descendingKesulitan() {
	var pass, idx, i int
	var temp project

	pass = 1
	for pass <= countProject-1 {
		idx = pass - 1
		i = pass

		for i < countProject {
			if daftarProject[idx].kesulitan > daftarProject[i].kesulitan {
				idx = i
			}
			i = i + 1
		}
		temp = daftarProject[pass-1]
		daftarProject[pass-1] = daftarProject[idx]
		daftarProject[idx] = temp

		pass = pass + 1
	}

}

func ascendingTanggal() {
	var pass, idx, i int
	var temp project
	var tanggalIdx, tanggalI int

	pass = 1
	for pass <= countProject-1 {
		idx = pass - 1
		i = pass

		for i < countProject {
			tanggalIdx = (daftarProject[idx].tahun * 360) + (daftarProject[idx].bulan * 30) + daftarProject[idx].day
			tanggalI = (daftarProject[i].tahun * 360) + (daftarProject[i].bulan * 30) + daftarProject[i].day

			if tanggalIdx > tanggalI {
				idx = i
			}
			i = i + 1
		}
		temp = daftarProject[pass-1]
		daftarProject[pass-1] = daftarProject[idx]
		daftarProject[idx] = temp

		pass = pass + 1
	}

}
func descendingTanggal() {
	var pass, idx, i int
	var temp project
	var tanggalIdx, tanggalI int

	pass = 1
	for pass <= countProject-1 {
		idx = pass - 1
		i = pass

		for i < countProject {
			tanggalIdx = (daftarProject[idx].tahun * 360) + (daftarProject[idx].bulan * 30) + daftarProject[idx].day
			tanggalI = (daftarProject[i].tahun * 360) + (daftarProject[i].bulan * 30) + daftarProject[i].day

			if tanggalIdx < tanggalI {
				idx = i
			}
			i = i + 1
		}
		temp = daftarProject[pass-1]
		daftarProject[pass-1] = daftarProject[idx]
		daftarProject[idx] = temp

		pass = pass + 1
	}

}

func statistikKategori() {
	if countProject == 0 {
		fmt.Println("Belum ada data kategori.")
		return
	}

	var categories [100]string
	var counts [100]int
	var catTotal int

	for i := 0; i < countProject; i++ {
		found := false
		for j := 0; j < catTotal; j++ {
			if categories[j] == daftarProject[i].kategori {
				counts[j]++
				found = true
				break
			}
		}
		if !found && catTotal < 100 {
			categories[catTotal] = daftarProject[i].kategori
			counts[catTotal] = 1
			catTotal++
		}
	}

	fmt.Println("Statistik Kategori:")
	for i := 0; i < catTotal; i++ {
		fmt.Printf("- %s: %d project\n", categories[i], counts[i])
	}
}

func keluar() {
	fmt.Println("Terima kasih telah menggunakan aplikasi!")
}

func main() {
	for {
		fmt.Println(".====================================.")
		fmt.Println("|         APLIKASI PORTOFOLIO        |")
		fmt.Println("'===================================='")
		fmt.Println("| 1. Tambah Project                  |")
		fmt.Println("| 2. Edit Project                    |")
		fmt.Println("| 3. Hapus Project                   |")
		fmt.Println("| 4. Keahlian                        |")
		fmt.Println("| 5. Cari Project                    |")
		fmt.Println("| 6. Urutkan Project                 |")
		fmt.Println("| 7. Statistik Kategori              |")
		fmt.Println("| 8. Keluar                          |")
		fmt.Println("'===================================='")
		fmt.Print("Pilih menu: ")

		var pilihan string
		fmt.Scan(&pilihan)

		switch pilihan {
		case "1":
			tambahProject()
		case "2":
			editProject()
		case "3":
			hapusProject()
		case "4":
			keahlian()
		case "5":
			cariProject()
		case "6":
			urutkanProject()
		case "7":
			statistikKategori()
		case "8":
			keluar()
			return
		default:
			fmt.Println("\nPilihan tidak valid!")
		}
	}
}
