package main

import "fmt"

func main() {
	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 无限循环，直到break/return
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		// 只打印奇数
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

// go run src/forloop/for.go
