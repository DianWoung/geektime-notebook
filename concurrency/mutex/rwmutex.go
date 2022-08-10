package mutex

import (
	"sync"
	"time"
)

//
//readers-writers 问题一般有三类，基于对读和写操作的优先级，读写锁的设计和实现也分成三类。
//Read-preferring：
//读优先的设计可以提供很高的并发性，但是，在竞争激烈的情况下可能会导致写饥饿。这是因为，如果有大量的读，这种设计会导致只有所有的读都释放了锁之后，写才可能获取到锁。
//Write-preferring：
//写优先的设计意味着，如果已经有一个 writer 在等待请求锁的话，它会阻止新来的请求锁的 reader 获取到锁，所以优先保障 writer。
//当然，如果有一些 reader 已经请求了锁的话，新请求的 writer 也会等待已经存在的 reader 都释放锁之后才能获取。
//所以，写优先级设计中的优先权是针对新来的请求而言的。这种设计主要避免了 writer 的饥饿问题。
//不指定优先级：
//这种设计比较简单，不区分 reader 和 writer 优先级，
//某些场景下这种不指定优先级的设计反而更有效，因为第一类优先级会导致写饥饿，第二类优先级可能会导致读饥饿，这种不指定优先级的访问不再区分读写，大家都是同一个优先级，解决了饥饿的问题。

// 如果你遇到可以明确区分 reader 和 writer goroutine 的场景，且有大量的并发读、少量的并发写，并且有强烈的性能需求，你就可以考虑使用读写锁 RWMutex 替换 Mutex
func main() {
	var counter CounterRW
	for i := 0; i < 10; i++ {
		go func() {
			for {
				counter.Count()
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for {
		counter.Incr()
		time.Sleep(time.Second)
	}
}

type CounterRW struct {
	mu    sync.RWMutex
	count uint64
}

//Incr use write lock
func (c *CounterRW) Incr() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

//Count use read lock
func (c *CounterRW) Count() uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}
