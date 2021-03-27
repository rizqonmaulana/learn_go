package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++{
		fmt.Println("Perulangan ke-", counter)
		
	}

	slice := []string{"Go", "Python", "Javascript"}
	for i := 0; i < len(slice); i++{
		fmt.Println(slice[i])
		
	}

	// for range
	// slice
	for index, name := range slice {
		fmt.Println("index i = ", index, " = ", name)
	}

	for _, name := range slice {
		fmt.Println("index i = ", name)
	}

	// map
	person := make(map[string]string)
	person["name"] = "Rizqon"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}