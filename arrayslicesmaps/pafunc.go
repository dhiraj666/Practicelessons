package main

import (
	"fmt"
)

func main() {

	var array [1e6]int
	// eight megabytes has to be allocated in stack everytime this function is called
	passArraytoFunc(array)
	// eight bytes needs to be copied only for passing that array to function
	//passArraytoFuncbyPointer(&array)
}

func passArraytoFunc(array [1e6]int) {
	fmt.Println(array)
}

/*func passArraytoFuncbyPointer(array *[1e6]int) {
	fmt.Println(array)
}*/
