package main

import "fmt"

/**
merupakan tipe data yang sangat spesial. Variabel bertipe ini bisa menampung segala jenis data, bahkan array, pointer, apapun. Tipe data dengan konsep ini biasa disebut dengan dynamic typing.
*/

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "upps"
	}
}

func main() {
	var data interface{} = Ups(2)
	fmt.Println(data)
}