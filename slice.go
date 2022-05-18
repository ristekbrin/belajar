package main

import (
	"fmt"
)

func main() {
	var months = [...]string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//months[5] = "Ganti"
	//fmt.Println(slice1)

	//slice1[1] = "Ganti Slice"
	//fmt.Println(months)

	//try append

	var slice2 = months[11:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Machrus")
	fmt.Println(slice3)
	slice3[1] = "Machrus Tohir"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	//membuat slice menggunakan make
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Machrus"
	newSlice[1] = "Tohir"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice from other slice
	copySlice := make([]string, 1, cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
