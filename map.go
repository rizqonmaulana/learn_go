package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Rizqon",
		"address": "Balikpapan",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["title"] = "Programmer"
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Go"
	book["author"] = "Rizqon"
	book["price"] = "123456"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "price")
	fmt.Println(book)
	fmt.Println(len(book))
		
}