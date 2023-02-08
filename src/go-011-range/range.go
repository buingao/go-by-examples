package main

import "fmt"

func main() {
	// use range to sum the numbers in a slice.
	nums := []int{2, 3, 4}
	sum := 0
	// range on arrays and slices provides both the index and value for each entry.
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// map key,value
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// Strings and Runes
	// range on strings iterates over Unicode code points
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
