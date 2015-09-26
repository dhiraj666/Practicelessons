package main

// slice can't be used as a key , we can use it as a map value but never as a key

import (
	"fmt"
)

func main() {

	var maps = make(map[string]string)

	mapone := map[int]string{1: "dhiraj", 2: "roshan"}

	fmt.Println(maps)

	value, exists := mapone[3]
	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("key doesn't exists", value)
	}

}
