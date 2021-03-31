package main

import "fmt"

// recover harus diletakan pada difer function, karena ketika panic terjadi maka kode dibawahnya tidak akan di eksekusi

func endApp()  {
	message := recover()

	if message != nil {
		fmt.Println("error dengan pesan : ", message)
	}
	
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool)  {
	defer endApp()

	if error {
		panic("Aplikasi Error")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(false)
}