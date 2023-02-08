package main

import "fmt"

/*
Buffered channels accept a limited number of values without a corresponding receiver for those values.
*/
func main() {
	// 如果不设置第二个参数，表示unbuffered
	// 运行报错，fatal error: all goroutines are asleep - deadlock!
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
