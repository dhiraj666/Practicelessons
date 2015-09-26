package main

import "fmt"

type arraystruct struct {
	a int
	b string
}

func main() {
	var arrayofint [5]int
	fmt.Println(arrayofint)

	arrayofint[0] = 2
	arrayofint[1] = 3
	arrayofint[4] = 5
	fmt.Println(arrayofint)

	as := [4]arraystruct{}
	fmt.Println(as[0], as[1], as[2], as[3])

	as[0] = arraystruct{
		2,
		"dhiraj",
	}

	fmt.Println(as[0])

	fmt.Println("use an array literal")

	arrayliteral := [5]int{1, 2, 4, 5, 6}
	fmt.Println(arrayliteral)

	fmt.Println("this is useful when we don't want to count then numbers of elements in array while declaring array literal")

	aldntel := [...]int{1, 2, 3, 4, 4, 5, 6, 7, 8, 9, 12, 34, 23}
	fmt.Println("print the arrayliteral with ...")
	fmt.Println(aldntel)

	fmt.Println("the number of elements is ", len(aldntel))

	arrayinit := [5]int{1: 2, 2: 3}

	fmt.Println(arrayinit)

	// we can have an array of pointers

	arraypt := [5]*int{new(int), new(int), new(int), new(int), new(int)}
	fmt.Println(arraypt)
	*arraypt[0] = 5
	fmt.Println(*arraypt[0])
	fmt.Println(arraypt)

	fmt.Println(*arraypt[3])

	// For one array to copy to another , the type and length should be same

	var da [5]int
	oa := [5]int{1, 2, 3, 4, 5}
	da = oa
	fmt.Println(oa)
	fmt.Println("destinationarray", da)

	//var desar [5]int
	//osar := [6]int{1, 2, 3, 4, 5, 6}
	//desar = osar
	//fmt.Println(desar, osar)

	// copying an array of pointers copies the pointer values and not the values that the pointer points to
}
