package main

import (
	"errors"
	"fmt"
)

//这意味着函数不但可以用于封装代码、分割功能、解耦逻辑，还可以化身为普通的值，
//在其他函数间传递、赋予变量、做类型判断和转换等等，就像切片和字典的值那样。
//而更深层次的含义就是：函数值可以由此成为能够被随意传播的独立逻辑组件（或者说功能模块）。

//只要两个函数的参数列表和结果列表中的元素顺序及其类型是一致的，我们就可以说它们是一样的函数，或者说是实现了同一个函数类型的函数。

type Printer func(contents string) (n int, err error)

func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

//高阶函数可以满足下面的两个条件：
//1. 接受其他的函数作为参数传入；
//2. 把其他的函数作为结果返回。
//只要满足了其中任意一个特点，我们就可以说这个函数是一个高阶函数。高阶函数也是函数式编程中的重要概念和特征。

type operate func(x, y int) int

type calculateFunc func(x int, y int) (int, error)

func genCalculator(op operate) calculateFunc {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	var p Printer
	p = printToStd
	i, err := p("hello world")
	if err != nil {
		return
	}
	fmt.Println(i)

	var x, y = 56, 78

	op := func(x, y int) int {
		return x * y
	}

	add := genCalculator(op)

	res, err := add(x, y)
	fmt.Printf("result is %d And err is %v\n", res, err)

}
