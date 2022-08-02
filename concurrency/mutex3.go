package main

func main() {
	//易错的四种场景
	// Lock/Unlock 不是成对出现， 如果mutex没有被加锁，这个时候解锁操作将会panic

	// Copy已使用的Mutex， mutex的state是公共的，如果复制这个对象之前已经是加锁状态，那么赋值之后就已经有锁了 go vet

	// 重入, Mutex是不可重入的锁. 可重入锁，又叫递归锁，

	// 死锁

}
