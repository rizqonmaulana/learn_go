package main

import "fmt"

func main() {
	//  break menghentikan semua proses looping setelahnya
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if(i == 5) {
			break
		}
	}

	// continue menyelesaikan looping saat itu dan langsung lanjut ke looping setelahnya
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}
}