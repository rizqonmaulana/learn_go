package main

import "fmt"

// untuk multiple return value perlu ditambahkan tipe data yang akan di return sesuai dengan jumlah return value
func getFullName() (string, string, string) {
	return "Rizqon", "Maulana", "Programmer"
}

func main() {
	// multiple return value
	firstName, lastName, title := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(title)

}