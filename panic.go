package main

import "fmt"

func endApp()  {
	fmt.Println("Aplikasi selesai")
}

// panic digunakan untuk sistem error jika ingin mengentikan aplikasi

func runApp(error bool)  {
	defer endApp()

	if error {
		panic("Aplikasi Error")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}