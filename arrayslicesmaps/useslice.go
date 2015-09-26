package main

import (
	"fmt"
)

func main() {

	// create a slice of integers
	sliceint := []int{10, 20, 30, 40, 50}
	fmt.Println(sliceint)

	newslice := sliceint[1:3]
	// For a slice[i:j] with an underlying array of capacity k
	// lenghth : j-i
	// capacity : k-i
	fmt.Println(newslice)
	newslice[0] = 25
	fmt.Println(sliceint)
	fmt.Println(newslice)
}
