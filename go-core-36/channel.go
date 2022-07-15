package main

import (
	"fmt"
	"sync"
)

//进入通道的并不是在接收操作符右边的那个元素值，而是它的副本。
// 它们处理元素值时都是一气呵成的，绝不会被打断。为了保证通道中元素值的完整性，也是为了保证通道操作的唯一性

//发送操作包括了“复制元素值”和“放置副本到通道内部”这两个步骤。
//接收操作通常包含了“复制通道内的元素值”“放置副本到接收方”“删掉原值”三个步骤。

//由于通道类型是引用类型，所以它的零值就是nil。换句话说，当我们只声明该类型的变量但没有用make函数对它进行初始化时，该变量的值就会是nil。我们一定不要忘记初始化通道！
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int, 100)
	wg := sync.WaitGroup{}
	go setMsg(ch1, ch2, &wg)
	go func() {
		wg.Add(1)
		for {
			i, ok := <-ch1
			if !ok {
				fmt.Println("ch1 is closed")
				break
			}
			ch3 <- i + 10000
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		for {
			x, ok := <-ch2
			if !ok {
				fmt.Println("ch2 is closed")
				break
			}
			ch3 <- x
		}
		wg.Done()
	}()
	wg.Wait()
	close(ch3)
	//遍历方式1，使用label来跳出循环
ch3Loop:
	for {
		select {
		case c, ok := <-ch3:
			if !ok {
				break ch3Loop
			}
			fmt.Println(c)
		default:
			fmt.Println("done")
		}
	}
	//遍历方式2，使用bool判断
	completed := false
	for !completed {
		select {
		case c, ok := <-ch3:
			if !ok {
				completed = true
			}
			fmt.Println(c)
		default:
			fmt.Println("done")
		}
	}
}

func setMsg(m chan<- int, n chan<- int, wg *sync.WaitGroup) {
	wg.Add(1)
	for i := 0; i < 10; i++ {
		if i > 3 {
			n <- i
		}
		m <- i
	}
	close(m)
	close(n)
	wg.Done()
}
