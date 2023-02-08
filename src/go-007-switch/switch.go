package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	/*
		You can use commas to separate multiple expressions in the same case statement. We use the optional default case in this example as well.
		可以使用逗号分隔同一个case语句中的多个表达式。在本例中，我们也使用可选的default case。
	*/
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	/*
		switch without an expression is an alternate way to express if/else logic.
		Here we also show how the case expressions can be non-constants.
		不带表达式的Switch是表达if/else逻辑的另一种方法。
		在这里，我们还将展示case表达式如何可以是非常量。
	*/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
		// 2023-02-07 16:43:00.655066 +0800 CST m=+0.000249714
		fmt.Println(t)
	}

	// 类型开关比较的是类型而不是值。您可以使用它来发现接口值的类型。在本例中，变量t的类型与其子句对应。
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

// go run src/switch/switch.go
