// for executable
package main

import (
	// for formatting
	"fmt"
	// for synchronization
	"sync"
	//
	"time"
)

var wg sync.WaitGroup

func main() {
	// create an unbuffered channel
	baton := make(chan int)
	// add a count of one for the last runner
	wg.Add(1)
	// First runner to his mark
	go Runner(baton)

	//start the race
	baton <- 1

	wg.Wait()

}

func Runner(baton chan int) {
	var newRunner int
	// wait to receive the baton
	runner := <-baton

	// start running around the track

	fmt.Printf("Runner %d Running with Baton\n", runner)

	// new runner to the line

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d to the line \n", newRunner)
		go Runner(baton)
	}

	// running around the track

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d Finished , run over \n", runner)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d Exchange with Runner %d\n", runner, newRunner)

	baton <- newRunner

}
