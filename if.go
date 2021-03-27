package main

import "fmt"

func main() {
	name := "maulana"

	if name == "rizqon" {
		fmt.Println("Hello " + name)
	} else if name == "maulana" {
		fmt.Println("Hello " + name)
	} else {
		fmt.Println("Kamu siapa ?")
	}

	// short statement if go
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}