package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 25
	return &p
}

func main() {
	// This syntax creates a new struct.
	fmt.Println(person{name: "gornin", age: 20})
	fmt.Println(person{"ben", 23})
	fmt.Println(person{name: "Alice"})
	// 前缀&产生指向结构体的指针
	fmt.Println(&person{"anna", 20})
	fmt.Println(newPerson("kevin"))

	/*
		Access struct fields with a dot.
		You can also use dots with struct pointers - the pointers are automatically dereferenced.
	*/
	s := person{"sean", 50}
	fmt.Println(s.name)
	s.age = 55
	fmt.Println(s)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 100
	fmt.Println(sp)
}
