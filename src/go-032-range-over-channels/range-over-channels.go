package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// Because we closed the channel above,
	// the iteration terminates after receiving the 2 elements.
	close(queue)
	// close(queue) // fatal error: all goroutines are asleep - deadlock!

	for elem := range queue {
		fmt.Println(elem)
	}
}
