package main

import (
	"fmt"
)

var slice = make([]int, 1e6)

func main() {
	slice = passslice(slice)
	fmt.Println(slice[0])
}

func passslice(slice []int) []int {
	return slice
}
