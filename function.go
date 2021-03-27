package main

import "fmt"

func sayHello(nama string)  {
	fmt.Println("Hello World", nama)
}

// untuk return value perlu ditambahkan tipe data yang akan di return
func getHello(nama string) string {
	return "Hello " + nama
}


func main() {
	nama := "adit"
	for i := 0; i < 10; i++ {
		sayHello(nama)
	}

	// function return value
	result := getHello("Rizqon")
	fmt.Println(result)

}