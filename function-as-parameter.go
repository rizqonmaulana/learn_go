package main

import "fmt"

type Filter func(string) string
//  baca params: menerima string & function, function memiliki paramater string, return dari function adalah string
func sayHelloWithFilter(name string, filter Filter) { 
	nameFiltered := filter(name)
	fmt.Print("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Fuck" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Fuck", spamFilter)
}