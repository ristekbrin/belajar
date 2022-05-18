package main

import "fmt"

func main() {
	//contoh break
	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan ke ", i)
		if i == 5 {
			break
		}
	}

	//contoh continue
	for i := 1; i < 10; i++ {

		if i%2 == 1 {
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}
