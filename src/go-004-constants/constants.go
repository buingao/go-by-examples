package main

import (
	"fmt"
	"math"
)

// const declares a constant value.
const s string = "constant"

func Constants() {
	fmt.Println(s)

	// A const statement can appear anywhere a var statement can.
	const n = 500000000
	// Constant expressions perform arithmetic with arbitrary precision.
	// 常量表达式以任意精度执行算术。
	const d = 3e20 / n
	fmt.Println(d)

	// A numeric constant has no type until it’s given one, such as by an explicit conversion.
	// 数值型常量在被赋予类型之前是没有类型的，比如通过显式转换。
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call.
	// For example, here math.Sin expects a float64.
	// 通过在需要类型的上下文中使用数字，例如变量赋值或函数调用，可以为数字指定类型。
	// 例如，这里的math.Sin期待一个float64类型的参数。
	fmt.Println(math.Sin(n))
}

func main() {
	Constants()
}

// go run go-by-examples/src/go-004-constants.go
