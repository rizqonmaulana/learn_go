package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {
	// anonymous function : membuat function tanpa membuat nama function
	// function bisa disimpan dalam var seperti dibawah
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("rizqon", blacklist)

	// atau anon func bisa ditaruh langsung sebagai params
	registerUser("root", func(name string) bool {
		return name == "root"
	})

	registerUser("rizqon", func(name string) bool {
		return name == "root"
	})
}