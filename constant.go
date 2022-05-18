package main

import "fmt"

func main() {
	const firstName string = "Machrus"
	const lastName = "Tohir"

	const value = 1000
	//membuat multiple constans deklrasi

	const (
		tes  = "Tester"
		data = "Data"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	fmt.Println(tes, data)
}
