package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == ""{
		return nil
	} else {
		return map[string]string {
			"name" : name,
		}
	}
}

func main() {
	// dengan nil
	var person map[string]string = NewMap("Rizqon")
	
	if person == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person)
	}

	// tanpa nil
	// person := map[string]string{
	// 	"name" : "Ayam",
	// }
	
	// if person["name"] == "" {
	// 	fmt.Println("Data Kosong")
	// } else {
	// 	fmt.Println(person)
	// }
}