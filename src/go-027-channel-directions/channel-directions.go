package main

import "fmt"

/*
使用通道作为函数参数时，可以指定通道是只发送值还是只接收值。
这种专一性增加了程序的类型安全性。
*/

/*
<-chan receive
chan<- send
*/

// This ping function only accepts a channel for sending values.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
