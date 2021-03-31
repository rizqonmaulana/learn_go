package main

import "fmt"

// jika menggunakan interface hasName maka harus memilii fungsi GetName yang memiliki hasil string
type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var rizqon Person
	rizqon.Name = "Rizqon"
	sayHello(rizqon)

	cat := Animal{
		Name : "Jiko",
	}

	sayHello(cat)
}