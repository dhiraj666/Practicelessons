package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// setting logical processor to one
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	fmt.Println("Starting go routines")
	go func() {

		defer wg.Done()

		for i := 0; i < 3; i++ {
			for c := 'A'; c < 'A'+26; c++ {
				fmt.Printf("go routine upperncase %c\n", c)
			}

		}

	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {

			for c := 'a'; c < 'a'+26; c++ {

				fmt.Printf(" go routine small case %c\n", c)
			}

		}

	}()
	wg.Wait()
	fmt.Println("terminating main program")

}
