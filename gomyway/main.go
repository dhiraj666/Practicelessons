package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {

	run()
	waitGroup.Wait()
}

func run() {

	waitGroup.Add(2000)
	for i := 0; i < 2000; i++ {
		go func() {
			fmt.Println("dhiraj")
			waitGroup.Done()
		}()

	}
	go func() {
		waitGroup.Wait()
	}()
}
