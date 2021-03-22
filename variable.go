package main 

import "fmt"

func main() {
	// penyebutan tipe data jika value variabel belum ditentukan
	var name string

	name = "Rizqon"
	fmt.Println(name)

	name = "Rizqon Maulana"
	fmt.Println(name)

// jika value variabel sudah ada maka penyebutan tipe data tidak wajib, go dapat mengenalinya
	var friend = "Amba"
	fmt.Println(friend)

	friend = "Baba"
	fmt.Println(friend)

	friend = "Baba Amba"
	fmt.Println(friend)


	var age = 20
	fmt.Println(age)

	age = 10
	fmt.Println(age)

	age = 45
	fmt.Println(age)

	// var tidak wajib, bisa diganti dengan := namun penggunaanya hanya untuk deklarasi awal
	country := "Indonesia"
	fmt.Println(country)

	country = "Belgium"
	fmt.Println(country)

	country = "Japan"
	fmt.Println(country)
	
	//  membuat lebih dari 1 variabel
	var (
		firstName = "Rizqon"
		lastName = "Maulana"
	)
	fmt.Println(firstName + " " + lastName)
	fmt.Println(lastName)
}