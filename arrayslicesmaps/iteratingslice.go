package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40}

	for index, value := range slice {
		fmt.Printf("Index : %d Value : %d \n", index, value)
	}

	for index, value := range slice {
		fmt.Printf("the element address : %X , value address : %X\n", &slice[index], &value)
	}
}

// len , cap works with
// arrays, slices and channels
// slice, channels and maps are reference types
