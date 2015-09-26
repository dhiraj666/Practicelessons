package main

import (
	"log"
	"os"
	"time"

	"github.com/dhiraj666/gocode/concurrencypatterns/runner"
)

// timeout is the number of second the program has to finish
const timeout = 3 * time.Second

// main is the entry point

func main() {
	log.Println("starting work")

	// create a new timer value for this run

	r := runner.New(timeout)

	// add the tasks to be run
	r.Add(createTask(), createTask(), createTask())

	// run the tasks and handle the result

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("terminating due to timeout")
			os.Exit(1)

		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt")
			os.Exit(2)
		}
	}
	log.Println("Process ended")

}

// createTask returns an example task that sleeps for the specified
//number of seconds based on the id

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
