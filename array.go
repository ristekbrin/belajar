package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Machrus"
	names[1] = "Tohir"
	names[2] = "Ganteng"

	fmt.Println(names[0], names[1], names[2])

	var values = [3]int{
		90,
		80,
		70,
	}
	fmt.Println(values)

	fmt.Println(len(names), len(values))
}
