package main

import "fmt"

func main() {
	counter := 0

	// ketika ada func child didalam func utama, maka func child dapat mengakses data func parent
	// counter di line 11 mengubah counter di line 6
	// jika counter di line 6 tidak mau diubah maka definisikan kembali counter didalam func increment
	increment := func() {
		// counter := 5
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
}