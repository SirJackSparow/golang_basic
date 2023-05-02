package main

import (
	"fmt"
	"math"
)

func main() {

	//make variable golang

	const (
		firstName = "Fendy"
		lastNAme  = "Maulana"
	)

	var (
		ages    = 24
		married = false
	)

	petName := "Ramzy"

	fmt.Println("Hello Worlds", math.Pi)

	fmt.Println(ages, married)

	fmt.Println(petName)

	// convertion

	var name = "Fendy"

	var e = name[0]

	var eString = string(e)

	fmt.Println(eString)

	//type data array

	var datax [3]string

	datax[0] = "bima"
	datax[1] = "maulana"
	datax[2] = "Effendy"

	var values = [2]int{11, 22}

	fmt.Println(datax[1])

	fmt.Println(values)

}

func condotion(n int8) int8 {
	if n > 30 {
		fmt.Println("Old")
	} else if n == 30 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Young")
	}

	return n
}

func forLooping(l int) {

	for i := 0; i < l; i++ {
		fmt.Println(i)
	}
}
