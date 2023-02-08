package main

import "fmt"

func main() {
	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements)
	// 与数组不同，slice只根据它们所包含的元素(而不是元素的数量)进行类型划分。
	// To create an empty slice with non-zero length, use the builtin make.
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	// len returns the length of the slice as expected
	fmt.Println("len:", len(s))

	// the builtin append, which returns a slice containing one or more new values
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be copyed.
	// Here we create an empty slice c of the same length as s
	c := make([]string, len(s))
	// and copy into c from s.
	copy(c, s)
	fmt.Println("cpy:", c)

	// index 2 3 4, not include index 5
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)

	// declare and initialize a variable for slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoDim := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDim[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDim[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoDim)
}
