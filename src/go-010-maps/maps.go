package main

import "fmt"

func main() {
	// make(map[key-type]val-type)
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m) // 形式：map[k:v k:v]

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m, len(m))

	// The optional second return value when getting a value from a map
	// indicates if the key was present in the map.
	// ignored it with the blank identifier _
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and initialize a new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
