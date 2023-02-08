package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {

	co := container{
		base: base{num: 1},
		str:  "some name",
	}
	// We can access the base’s fields directly on co
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	// Alternatively, we can spell out the full path using the embedded type name.
	fmt.Println("also num:", co.base.num)
	// invoke a method that was embedded from base directly on co.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}
	// a container now implements the describer interface
	// because it embeds base.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
