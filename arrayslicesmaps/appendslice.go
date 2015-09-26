package main

import (
	"fmt"
)

func main() {
	orgslice := []int{10, 20, 30, 40, 50}
	fmt.Println(orgslice)

	newslice := orgslice[1:3]

	fmt.Println(newslice)

	newslice = append(newslice, 60)
	fmt.Println("after appending", newslice)

	fmt.Println("originalslice", orgslice)

	newslice = append(newslice, 70)
	fmt.Println("reached the capacity of newslcie", newslice)
	newslice = append(newslice, 80)
	fmt.Println("newslice when capacity got over", newslice)
	fmt.Println("what happened to orgslice", orgslice)
	fmt.Println("length of newslice", len(newslice))
	fmt.Println("Capcity of newslice", cap(newslice))
	fmt.Println("length of original slice:", len(orgslice))
	fmt.Println("capacity of original slice ", cap(orgslice))

	sliceone := []int{10, 20, 30, 40, 50}
	newsliceone := append(sliceone, 60)
	fmt.Println("sliceone values :", sliceone)
	fmt.Println("new sliceone values after appending:", newsliceone)
	fmt.Println(len(sliceone), len(newsliceone), cap(sliceone), cap(newsliceone))

	// three indexes slices option

	sourcesl := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	tis := sourcesl[2:3:4]

	fmt.Println("sourceslice", sourcesl)
	fmt.Println("thirdindexslice", tis)

}
