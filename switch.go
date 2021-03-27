package main

import "fmt"

func main() {
	name := "Python999"

	switch name {
	case "Rizqon":
		fmt.Println("Hallo Rizqon")
	case "Maulana":
		fmt.Println("Hallo Maulana")
	case "Go":
		fmt.Println("Hallo Go")
	case "Python":
		fmt.Println("Hallo Python")
	default : 
		fmt.Println("Kamu siapa ?")
	}

	// short statement
	// switch length := len(name); length > 5 {
	// case true :
	// 	fmt.Println("Nama Terlalu panjang")
	// case false :
	// 	fmt.Println("Nama sudah benar")
	// }

	// switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu panjang")
	case length > 8:
		fmt.Println("Nama Terlalu panjang, kurangin dikit")
	case length > 6:
		fmt.Println("Nama Sudah pas")
	case length < 4:
		fmt.Println("Nama kurang panjang")

		} 
}