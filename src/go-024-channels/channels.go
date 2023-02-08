package main

import "fmt"

/*
By default sends and receives block until both the sender and receiver are ready.
*/
func main() {
	// Create a new channel with make(chan val-type)
	messages := make(chan string)

	go func() {
		// Send a value into a channel using the channel <- syntax.
		messages <- "ping"
	}()

	// The <-channel syntax receives a value from the channel.
	msg := <-messages
	fmt.Println(msg)
}
