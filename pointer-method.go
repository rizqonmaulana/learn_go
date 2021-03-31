package main

import "fmt"

type Man struct {
	Name string
}

// tambahkan * pada stuct untuk membuatnya menjadi pointer method
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	rizqon := Man{"rizqon"}
	rizqon.Married()

	fmt.Println(rizqon.Name)
}