/*
在Go中，通过显式的、单独的返回值来传达错误是惯用的。
这与Java和Ruby等语言中使用的异常以及c中有时使用的重载单一结果/错误值形成了对比。
Go的方法使得很容易看到哪些函数返回错误，并使用用于任何其他非错误任务的相同语言结构来处理它们。
*/

package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int    // 错误值
	prob string // 错误描述
}

// 定义一个方法
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// 在本例中，我们使用&argError语法来构建一个新的结构体，为arg和prob两个字段提供值。
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	/*
		下面的两个循环测试了我们的每个错误返回函数。
		请注意，在if行上使用内联错误检查是Go代码中的常见习惯用法。
	*/
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	/*
		如果希望以编程方式使用自定义错误中的数据，则需要通过 类型断言 获取错误作为自定义错误类型的实例。
	*/
	_, e := f2(42)
	fmt.Println(e) // 42 - can't work with it
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
