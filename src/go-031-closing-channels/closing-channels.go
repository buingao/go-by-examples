package main

import "fmt"

/*
关闭通道表示不再在该通道上发送值。这对于将完成信息传递给通道的接收者非常有用。
*/
func main() {
	/*
		在本例中，我们将使用job通道将main()goroutine 要完成的工作传递给worker goroutine。
		当我们没有更多的工作给worker时，我们将关闭jobs channel。
	*/
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 接受 worker
	go func() {
		for {
			// the more value will be false if jobs has been closed
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
