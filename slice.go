package main

import "fmt"

func main() {
	// data slice, slice can change their sizes, slice is slice of array
	// ... if dont know size
	var dataArray = [...]string{
		"Jakarta",
		"Bandung",
		"Surabaya",
		"Batam",
		"Batavia",
		"Lain",
	}

	var slice1 = dataArray[1:3]

	fmt.Println(slice1)

	fmt.Println("Length", len(slice1))            // will call 2 cause [1:3] in data Bandung and Surabaya
	fmt.Println("Capacity of Slice", cap(slice1)) // the casity of slice start Bandung until Lain the size should be 5 [1:3] the size would start bandung until lain
}
