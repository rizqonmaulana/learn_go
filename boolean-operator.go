package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 55

	var lulusUjian = ujian >=75
	var lulusAbsensi = absensi >= 60

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)

	lulus = lulusUjian || lulusAbsensi
	fmt.Println(lulus)
}