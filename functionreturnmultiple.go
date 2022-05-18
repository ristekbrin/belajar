package main

import "fmt"

func getFullName() (string, string) {
	return "Machrus", "Tohir"
}
func main() {
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
