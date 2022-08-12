package main

import (
	"fmt"
	"sync"
)

var A int

func main() {
	var o sync.Once

	f1 := func() {
		fmt.Println("in f1")
		A = 1
	}

	o.Do(f1)

	f2 := func() {
		A = 2
	}

	o.Do(f2)
	fmt.Println(A)
}
