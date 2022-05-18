package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpMachrus NoKTP = "1212121212121"
	var marriedStatus Married = true
	fmt.Println(noKtpMachrus)
	fmt.Println(marriedStatus)
}
