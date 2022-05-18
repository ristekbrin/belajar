package main

import "fmt"

func main() {
	name := "yuna"

	switch name {
	case "Machrus":
		fmt.Println("Hallo Machrus Tohir")
	case "Yuna":
		fmt.Println("Hallo Yuna")
	default:
		fmt.Println("Hi, Kenalayan yuk")
	}

	/*switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}*/

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("Nama terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama sudah benar")

	}

}
