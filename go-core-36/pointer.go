package main

import (
	"fmt"
	"unsafe"
)

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}
func main() {
	// 示例1。
	dog := Dog{"little pig"}
	dogP := &dog
	dogPtr := uintptr(unsafe.Pointer(dogP))

	namePtr := dogPtr + unsafe.Offsetof(dogP.name)
	nameP := (*string)(unsafe.Pointer(namePtr))
	fmt.Printf("nameP == &(dogP.name)? %v\n",
		nameP == &(dogP.name))
	fmt.Printf("The name of dog is %q.\n", *nameP)

	*nameP = "monster"
	fmt.Printf("The name of dog is %q.\n", dogP.name)
	fmt.Println()

	// 示例2。
	// 下面这种不匹配的转换虽然不会引发panic，但是其结果往往不符合预期。
	numP := (*int)(unsafe.Pointer(namePtr))
	num := *numP
	fmt.Printf("This is an unexpected number: %d\n", num)

	//unsafe.Pointer可以表示任何指向可寻址的值的指针，同时它也是前面提到的指针值和uintptr值之间的桥梁。
	//也就是说，通过它，我们可以在这两种值之上进行双向的转换。
	//Go 语言中的哪些值是不可寻址的吗？
	//常量的值。
	//基本类型值的字面量。
	//算术操作的结果值。
	//对各种字面量的索引表达式和切片表达式的结果值。
	//不过有一个例外，对切片字面量的索引结果值却是可寻址的。
	//对字符串变量的索引表达式和切片表达式的结果值。
	//对字典变量的索引表达式的结果值。
	//函数字面量和方法字面量，以及对它们的调用表达式的结果值。
	//结构体字面量的字段值，也就是对结构体字面量的选择表达式的结果值。
	//类型转换表达式的结果值。
	//类型断言表达式的结果值。
	//接收表达式的结果值。

	//不可寻址的属性：
	//不可变：不可变的值不可寻址。常量、基本类型的值字面量、字符串变量的值、函数以及方法的字面量都是如此。其实这样规定也有安全性方面的考虑。绝大多数被视为临时结果的值都是不可寻址的。
	//临时结果：算术操作的结果值属于临时结果，针对值字面量的表达式结果值也属于临时结果。但有一个例外，对切片字面量的索引结果值虽然也属于临时结果，但却是可寻址的。
	//不安全的： 若拿到某值的指针可能会破坏程序的一致性，那么就是不安全的，该值就不可寻址。由于字典的内部机制，对字典的索引结果值的取址操作都是不安全的。另外，获取由字面量或标识符代表的函数或方法的地址显然也是不安全的。
}
