package main

import "fmt"

// named return value | jika tipe data return value sama maka penyebutan tipe data bisa hanya 1x
func getFullBio() (namaDepan, namaBelakang, job string) {
	namaDepan = "Rizqon"
	namaBelakang = "Maulana"
	job = "Software Engineer"

	// return tidak wajib memanggil nama return karena sudah di deklarasikan diatas, tapi jika mau panggil lagi boleh
	return
}

func main() {

	// beri underscore jika tidak mau mengambil return value, sesuaikan dengan urutan return
	namaAwal, _, pekerjaan := getFullBio()
	fmt.Println(namaAwal)
	fmt.Println(pekerjaan)

	// named return value, nama variabel disini tidak harus sama dengan return variabel diatas
	namaDepan, namaBelakang, job := getFullBio()
	fmt.Println(namaDepan, namaBelakang, job)
}