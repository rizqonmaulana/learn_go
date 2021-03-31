package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// implementasi pointer di function
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {

	// tambahkan & sebelum memanggil variabel
	// agar go mengambil data by reference
	// karena pada defaultnya go akan menduplikasi value variabel
	// dan tidak akan mengubah value variabel asal
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address3 := &address1

	address2.City = "Bandung"
	/**
		operator * digunakan untuk memaksa agar semua 
		yang mengakses struct yang sama akan mengikuti value nya	
	*/
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// new digunakan untuk membuat data baru yang kosong tanpa mengikuti value dari yang lain
	address4 := new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address{
		City: "Subang",
		Province: "Jawa Barat",
		Country: "",
	}

	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)

	
}