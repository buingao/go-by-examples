package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plus2(a, b, c int) int {
	return a + b + c
}

// multiple return values
func vals() (int, int) {
	return 3, 7
}

// 可变函数
// can be called with any number of trailing arguments
func sum(nums ...int) {
	fmt.Print(nums, "  ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println("sum:", total)
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plus2(2, 3, 4)
	fmt.Println("2+3+4=", res)

	// multiple assignment
	a, b := vals()
	fmt.Println(a, b)

	// only want a subset of the returned values
	_, c := vals()
	fmt.Println(c)

	// Variadic Functions 可变的函数
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	// func(slice...)
	sum(nums...)
}
