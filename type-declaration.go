package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpRizqon NoKTP = "1234455667774533"
	var marriedStatus Married = false

	fmt.Println(noKtpRizqon)
	fmt.Println(marriedStatus)
}