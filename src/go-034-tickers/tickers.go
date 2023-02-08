package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			// use the select builtin on the channel to await the values as they arrive every 500ms
			select {
			case <-done:
				return
			case t := <-ticker.C: // a channel that is sent values
				fmt.Println("tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
