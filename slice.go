package main

import "fmt"


func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var  slice1 = months[0:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))


	months[5] = "Changed"
	fmt.Println(slice1)

	slice1[0] = "new month"
	fmt.Println(months)

	var slice2 = months[9:12]
	fmt.Println(slice2)
	fmt.Println(cap(slice2))

	// append akan mengubah array sebelumnya jika capacity 
	// masih cukup, tapi jika capacity array telah penuh maka akan membuat array baru

	var slice3 = append(slice2, "Rizqon")
	fmt.Println(slice3)

	// cara membuat slice paling aman, menyembunyikan array didalam slice
	newSlice := make([]string,2,5)

	newSlice[0] = "Rizqon"
	newSlice[1] = "Maulana"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice (pastikan ukurannya sama, jika tidak akan terpotong)
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
	
	
	// jika tidak ditentukan panjang array maka jadi slice
	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)


}