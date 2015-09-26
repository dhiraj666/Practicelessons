// Package runner manages the running and lifetime of a process
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner runs a set of tasks within a given timeout and can be
// shutdown on an operating system interrupt
type Runner struct {
	//interrupt channel reports a signal from the operating system
	interrupt chan os.Signal

	//complete channel reports that processing is done
	complete chan error

	// timeourt reports that time has run out
	timeout <-chan time.Time

	//tasks holds a set of functions that are executed synchronously in index order
	tasks []func(int)
}

// ErrTimeout is returned when a value is received on the time out
var ErrTimeout = errors.New("received interrupt")

// ErrInterrupt is returned when an event from the os is received
var ErrInterrupt = errors.New("received interrupt")

// New returns a new ready to use runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

//Add attaches a task to the runner. Task is a function that takes an int ID
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start runs all the tasks and monitors channel events

func (r *Runner) Start() error {
	// we want to receive all interrupt based signals
	signal.Notify(r.interrupt, os.Interrupt)

	// Run the different tasks on a different goroutines
	go func() {
		r.complete <- r.run()
	}()

	select {
	// signaled when processing is done
	case err := <-r.complete:
		return err

		// signaled when we run out of time
	case <-r.timeout:
		return ErrTimeout
	}

}

// run execute each registered task

func (r *Runner) run() error {

	for id, task := range r.tasks {
		//Check for an interrupt signal from the OS
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		// execute the registered task
		task(id)
	}
	return nil
}

// gotInterrupt verifies if the interrupt signal has been issued
func (r *Runner) gotInterrupt() bool {
	select {
	//Signaled when an interrupt event is sent
	case <-r.interrupt:
		//Stop receiving any further signals
		signal.Stop(r.interrupt)
		return true
	// continue running as normal
	default:
		return false
	}
}
