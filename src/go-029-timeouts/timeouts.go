/*
超时对于连接到外部资源或需要限制执行时间的程序很重要。
由于通道和选择，在Go中实现超时非常简单和优雅。
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	// make 一个channel，带缓冲
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		// <-time.After awaits a value to be sent after the timeout of 1s
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
