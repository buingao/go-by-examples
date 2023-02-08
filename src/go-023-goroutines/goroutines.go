package main

import (
	"fmt"

	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// the usual way, running it synchronously
	f("direct")

	// will execute concurrently with the calling one
	go f("goroutine")

	// start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going~")

	time.Sleep(time.Second)
	fmt.Println("done")
}
