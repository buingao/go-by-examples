package main

import "fmt"

// This function intSeq returns another function
func intSeq() func() int {
	i := 0
	return func() int {
		// The returned function closes over the variable i to form a closure
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
