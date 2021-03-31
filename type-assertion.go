package main

import "fmt"

// type assertion digunakan untuk mengubah tipe data
// fitur ini sering digunakan ketika bertemu dgn data interface kosong

// func dgn return interface kosong
func random() interface{} {
	return 123
}

// type assertion 
func main() {
	var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// jika menggunakan type assertion lebih aman menggunakan
	// switch case untuk menghindari panic
	switch value := result.(type) {
		case string:
			fmt.Println("value", value, "is string")
		case int:
			fmt.Println("value", value, "is int")
		default:
			fmt.Println("unknown")
	}
}