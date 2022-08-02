package main

import (
	"fmt"
	"sync"
)

// go vet 可以检测copy lock问题
type Counter3 struct {
	sync.Mutex
	Count int
}

func main() {
	var c Counter3
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c) // 复制锁
}

// 这里Counter的参数是通过复制的方式传入的
func foo(c Counter3) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}
