// buffered channels

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // number of goroutines to use
	taskLoad         = 10 // amount of work to process
)

// wg is used to wait for the program to finish
var wg sync.WaitGroup

// init is called  so that
// go runtime can use it

func init() {
	rand.Seed(time.Now().Unix())
}

// main is the entry point

func main() {
	// create a buffered channel to manage the taskload

	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)
	// launch goroutines to handle the work
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// add a bunch of work to get done

	for post := 0; post < taskLoad; post++ {
		tasks <- fmt.Sprintf("Task :  %d", post)
	}

	// close the channel so the goroutines will quit
	// when the work is done
	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		// wait for thw work to be assigned

		task, ok := <-tasks
		if !ok {
			// this means the channel is empty and closed
			fmt.Printf("Worker : %d : shutting down\n", worker)
			return
		}
		//display we are starting the work
		fmt.Printf("worker : %d : started %s\n", worker, task)

		// randomly wait to simulate work time

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		// display we finshed the work
		fmt.Printf("worker: %d : completed %s\n", worker, task)

	}

}
