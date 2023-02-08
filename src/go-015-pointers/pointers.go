package main

import "fmt"

func main() {
	initial := 1
	fmt.Println("initial:", initial)

	zeroval(initial)
	fmt.Println("zeroval:", initial)

	// The &i syntax gives the memory address of initial
	zeroptr(&initial)
	fmt.Println("zeroptr:", initial)
}

// zeroval will get a copy of ival distinct from the one in the calling function.
func zeroval(val int) {
	val = 0
}

// zeroptr in contrast has an *int parameter, meaning that it takes an int pointer.
func zeroptr(ptr *int) {
	// The *iptr code in the function body then dereferences the pointer from its memory address to the current value at that address.
	// Assigning a value to a dereferenced pointer changes the value at the referenced address.
	*ptr = 0
}

// 内存地址 赋值给 指针变量
// var ptr *int = &initial
