/*
slices are three field
data structures
pointer to the underlying array
length and capacity
*/

package main

import (
	"fmt"
)

func main() {
	// create a slice of strings
	// that will contain the length and capacity of 5 elements

	slice := make([]string, 5)
	fmt.Println(slice)

	sliceint := make([]int, 2, 4)
	fmt.Println(sliceint)
	fmt.Println(sliceint[1])

	// sliceliterals are the idiomatic way

	slicelit := []int{1, 2, 3, 4, 5}
	slicelitst := []string{"dhiraj", "das", "rakesh", "das", "shyam", "das", "puja", "das", "shilpi", "das"}
	fmt.Println(slicelit)
	fmt.Println(slicelitst)

	// Initialize the index that represents the length and capacity we need
	// sliceliteral

	sliceliteralc := []int{99: 2}
	fmt.Println(len(sliceliteralc))

	var nilslice []int
	fmt.Println(nilslice)
	emptyslice := []int{}
	fmt.Println(emptyslice)
}
