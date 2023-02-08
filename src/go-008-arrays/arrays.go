package main

import "fmt"

/*
In typical Go code, slices are much more common;在典型的Go代码中，切片更为常见
arrays are useful in some special scenarios.数组在某些特殊场景中很有用。
*/
func main() {
	var a [5]int
	fmt.Println("a:", a) // [0 0 0 0 0]

	var s [5]string
	fmt.Println("s:", s) // [    ]

	a[4] = 100
	fmt.Println("set:", a)
	// get a value with array[index]
	fmt.Println("get:", a[4])
	// The builtin len returns the length of an array.
	fmt.Println("len:", len(a))

	// Use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	// compose types to build multi-dimensional data structures
	var twoDim [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoDim[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoDim) // [[0 1 2] [1 2 3]]
}

// go run src/arrays/arrays.go
