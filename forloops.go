package main

import "fmt"

func main() {
	/*counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}*/

	//for loops using statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	slice := []string{"Machrus", "Tohir", "Ganteng", "Sekali"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//coba for range
	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Machrus"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
