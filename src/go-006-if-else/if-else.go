package main

import "fmt"

func main() {
	// There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
	// 在Go中没有三元if，所以即使对于基本条件，也需要使用完整的if语句。
	if 7%2 == 0 {
		fmt.Println("7 is even") // 偶数
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 语句可以放在条件句前面；在此语句中声明的任何变量都可以在当前和所有后续分支中使用。
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
