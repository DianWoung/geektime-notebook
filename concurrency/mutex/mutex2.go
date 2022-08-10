package mutex

import (
	"fmt"
	"sync"
)

func main() {
	var counter Counter1

	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				counter.Incr()
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter.Count())
}

// Counter1 线程安全的计数器类型
type Counter1 struct {
	CounterType int
	Name        string

	mu    sync.Mutex
	count uint64
}

func (c *Counter1) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

//为什么需要对读取操作也加锁呢？
//因为写操作并不是原子的，一条普通的赋值语句其实并不是一个原子操作。
//比如在 32位机器上，写 int64 类型的变量就会有中间状态，因为它会被拆成两次 MOV 操作 -- 写低 32 位 和高 32 位。
//如果一个线程刚写完低 32 位，还没来得及写 高 32 位时，另一个线程读取了这个变量，那么就会得到一个毫无意义的中间变量，这可能使我们的程序出现诡异的 Bug。
//所以 sync/atomic 提供了对基础类型的一些原子操作，比如 LoadX, StoreX, SwapX, AddX，CompareAndSwapX 等。
//这些操作在不同平台有不同的实现，比如 LoadInt64 在 amd64 下就是一条简单的加载，但是在 386 平台下就需要更复杂的实现。

func (c *Counter1) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
