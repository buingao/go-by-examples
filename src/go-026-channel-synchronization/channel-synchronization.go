package main

import (
	"fmt"
	"time"
)

// in another goroutine
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// Send a value to notify that we’re done
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	// If you removed the <- done line from this program,
	// the program would exit before the worker even started.
	<-done
}
