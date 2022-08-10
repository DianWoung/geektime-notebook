package main

import (
	"fmt"
	"sync"
	"time"
)

//关于如何避免错误使用 WaitGroup 的情况，我们只需要尽量保证下面 5 点就可以了：
//不重用 WaitGroup。新建一个 WaitGroup 不会带来多大的资源开销，重用反而更容易出错。
//保证所有的 Add 方法调用都在 Wait 之前。
//不传递负数给 Add 方法，只通过 Done 来给计数值减 1。
//不做多余的 Done 方法调用，保证 Add 的计数值和 Done 方法调用的数量是一样的。
//不遗漏 Done 方法的调用，否则会导致 Wait hang 住无法返回。

// 线程安全的计数器
type Counter struct {
	mu    sync.Mutex
	count uint64
}

// 对计数值加一
func (c *Counter) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// 获取当前的计数值
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// sleep 1秒，然后计数值加1
func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}

func main() {
	var counter Counter

	var wg sync.WaitGroup
	wg.Add(10) // WaitGroup的值设置为10

	for i := 0; i < 10; i++ { // 启动10个goroutine执行加1任务
		go worker(&counter, &wg)
	}
	// 检查点，等待goroutine都完成任务
	wg.Wait()
	// 输出当前计数器的值
	fmt.Println(counter.Count())
}
