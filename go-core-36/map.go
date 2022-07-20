package main

import "fmt"

//1,键的类型必须是可以施加判等操作符的!=和==,函数，字典，切片，都不支持作为键
//2,求hash和判等操作的速度越快，对应的类型就越适合作为键类型
//3,优先选用数值类型和指针类型，通常情况下类型的宽度越小越好。如果非要选择字符串类型的话，最好对键值的长度进行额外的约束。

func removeMap(m map[int]string, n int) {
	delete(m, n)
}

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

	//所有传给函数的参数值都会被复制，函数在其内部使用的并不是参数值的原值，而是它的副本。由于数组是值类型，所以每一次复制都会拷贝它，以及它的所有元素值。
	//我在modify函数中修改的只是原数组的副本而已，并不会对原数组造成任何影响。注意，对于引用类型，比如：切片、字典、通道，像上面那样复制它们的值，只会拷贝它们本身而已，并不会拷贝它们引用的底层数据。
	//也就是说，这时只是浅表复制，而不是深层复制。以切片值为例，如此复制的时候，只是拷贝了它指向底层数组中某一个元素的指针，以及它的长度值和容量值，而它的底层数组并不会被拷贝。

	m := make(map[int]string, 0)
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"
	fmt.Println(m)
	var n = make(map[int]string, 0)
	n = m
	var o = make(map[int]string, 0)
	o = n
	fmt.Println(o)
	removeMap(o, 2)
	delete(o, 1)
	fmt.Println("=====")
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
}
