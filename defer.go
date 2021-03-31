package main

import "fmt"

func logging()  {
	fmt.Println("Selesai memanggil function")
}

// defer : digunakan untuk menjalankan func setelah func tertentu selesai dijalankan
// defer ditaruh pada line awal func agar tetap berjalan walaupun kode dibawahnya error

func runApplication(value int)  {
	defer logging()
	result := 10 / value
	fmt.Println("result", result)
	fmt.Println("Run application")
}

func main() {
	runApplication(0)
}