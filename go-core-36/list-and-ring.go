package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Node struct {
	Id   int
	Name string
}

//list链表
func getList() *list.List {
	rand.Seed(time.Now().UnixNano())
	l := list.New()
	for i := 0; i < 10; i++ {
		l.PushBack(&Node{
			Id:   i,
			Name: strconv.Itoa(rand.Intn(10)),
		})
	}
	return l
}

//ring环形链表
func getRing() *ring.Ring {
	rand.Seed(time.Now().UnixNano())
	r := ring.New(10)
	for i := 0; i < 10; i++ {
		r.Value = &Node{
			Id:   i,
			Name: strconv.Itoa(rand.Intn(10)),
		}
		r = r.Next()
	}
	return r
}

func main() {
	fmt.Println("+====start=====")
	//链表遍历方式
	var l *list.List
	l = getList()
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println("list's value is ", e.Value.(*Node).Id)
	}

	//环形链表遍历方式
	var r *ring.Ring
	r = getRing()
	//用do来遍历
	r.Do(func(n any) {
		fmt.Println("ring's value is ", n.(*Node).Id)
	})
	fmt.Println("===end===")
}

//
//语句var l list.List声明的链表l可以直接使用，这是怎么做到的呢？
//关键在于它的“延迟初始化”机制。所谓的延迟初始化，你可以理解为把初始化操作延后,
//仅在实际需要的时候才进行。延迟初始化的优点在于“延后”，它可以分散初始化操作带来的计算量和存储空间消耗。
//例如，如果我们需要集中声明非常多的大容量切片的话，那么那时的 CPU 和内存空间的使用量肯定都会一个激增，
//并且只有设法让其中的切片及其底层数组被回收，内存使用量才会有所降低。
