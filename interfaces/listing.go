package main

import (
	"fmt"
)

type duration int

func (d duration) pretty() string {
	return fmt.Sprintf("Duration : %d", d)
}

func main() {
	//var dur duration = 42
	//fmt.Println(dur.pretty())

	duration(42).pretty()
}
