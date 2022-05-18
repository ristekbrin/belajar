package main

import "fmt"

func main() {

	var person = map[string]string{
		"name":    "Machrus Tohir",
		"Address": "Malang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["Address"])

	var book map[string]string = make(map[string]string)
	book["tile"] = "Belajar Go-Lang"
	book["author"] = "Machrus Tohir"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
