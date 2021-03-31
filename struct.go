package main

import "fmt"

// struct = template data, struct tidak bisa langsung digunakan

type Customer struct {
	Name string
	Address string
	Age int
}

func sayHello(customer Customer, name string) {
	fmt.Println("Hello",name,"My name is", customer.Name)
}

// struct method
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello",name,"My name is", customer.Name)
}

func main() {
	var nana Customer
	nana.Name = "nana"
	nana.Address = "Balikpapan"
	nana.Age = 44
	
	fmt.Println(nana.Name)
	fmt.Println(nana.Address)
	fmt.Println(nana.Age)

	// sayHello(nana, "Joko")

	nana.sayHi("Joko")

	// struct literal

	// lulu := Customer{
	// 	Name : "lulu lala",
	// 	Address : "Bekasi",
	// 	Age : 99,
	// } 

	// fmt.Println(lulu)

	// struct function

}