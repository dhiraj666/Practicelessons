// Sample program to demonstrate
// an unbuffered channel to simulate a game of tennis between two go routines

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg is used to wait for the program to finish

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// create an unbuffered channel
	court := make(chan int)
	// add a count of two , one for each goroutine
	wg.Add(2)

	// launch two players
	go player("Nadal", court)
	go player("Djokovic", court)

	// start the set
	court <- 1

	wg.Wait()

}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("Player %s won\n", name)
			return
		}

		// pick a random number and see if we miss the ball

		n := rand.Intn(100)
		if n%2 == 0 {
			fmt.Printf("Player %s missed\n", name)

			close(court)
			return
		}

		fmt.Printf("Player %s hit %d\n", name, ball)
		ball++
		// hit the ball back to the opposition player
		court <- ball
	}

}
