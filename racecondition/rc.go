/*
when two or more go routines have unsynchronized
access to a shared resource and attempt to read
or write to that resource at the same time , we
have what's called a race condition
Read and write operations against a shared resource
must always be atomic
done by one goroutine at a time
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int // variable incremented by all goroutines

	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final counter :", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for i := 0; i < 2; i++ {
		// capture the value of counter
		value := counter

		// yield the thread and be placed back in queue
		runtime.Gosched()
		value++

		counter = value

	}

}
