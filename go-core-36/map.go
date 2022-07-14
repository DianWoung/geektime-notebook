package main

import "fmt"

//1,键的类型必须是可以施加判等操作符的!=和==,函数，字典，切片，都不支持作为键
//2,求hash和判等操作的速度越快，对应的类型就越适合作为键类型
//3,优先选用数值类型和指针类型，通常情况下类型的宽度越小越好。如果非要选择字符串类型的话，最好对键值的长度进行额外的约束。

func main() {
	aMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	k := "two"
	v, ok := aMap[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not found!")
	}
}
