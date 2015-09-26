package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Create go routines")
	go printPrime("A")
	go printPrime("B")

	wg.Wait()
	fmt.Println("terminating the main thread now ")

}

func printPrime(prefix string) {
	// schedule to call to Done to tell main we are done
	defer wg.Done()
next:
	for i := 0; i < 5000; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue next
			}
			fmt.Printf("%s:%d\n", prefix, i)
		}
	}
	fmt.Println("completed", prefix)
}
