package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// 如果在messages上有一个值，那么select将使用<-messages case作为该值
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 这里msg不能发送到消息通道，因为通道没有缓冲区，也没有接收器。因此选择默认情况。
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
