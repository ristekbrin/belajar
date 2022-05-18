package main

import "fmt"

func main() {

	var name = "yuna"

	if name == "Machrus" {
		fmt.Println("Hallo Machrus Tohir")
	} else if name == "yuna" {
		fmt.Println("Halo Yuna")
	} else {
		fmt.Println("Hi, Kenalan donk")
	}

	//short statement if in golang

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
