package main

import "fmt"

func main() {
	var name1 = "rizqon"
	var name2 = "rizqon"

	var result = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	result = value1 > value2
	fmt.Println(result)

	result = value2 > value1
	fmt.Println(result)
}