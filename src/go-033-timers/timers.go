package main

import (
	"fmt"
	"time"
)

func main() {
	// 计时器表示未来的单个事件。你告诉计时器你想要等待多长时间，它会提供一个通道在那个时候被通知。
	// 这个计时器将等待2秒。
	timer1 := time.NewTimer(2 * time.Second)

	// <-timer1.C阻塞在计时器的通道C上，直到通道C发送一个值指示计时器已触发。
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
