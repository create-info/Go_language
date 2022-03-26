package sync_primitives

//初版的互斥锁

//通过一个 flag 变量，标记当前的锁是否被某个 goroutine 持有。如果这个 flag 的值是 1，就代表锁已经被持有，那么，其它竞争的 goroutine 只能等待；如果这个 flag 的值是 0，就可以通过 CAS（compare-and-swap，或者 compare-and-set）将这个 flag 设置为 1，标识锁被当前的这个 goroutine 持有了。

//CAS 指令将给定的值和一个内存地址中的值进行比较
//CAS 是实现互斥锁和同步原语的基础，我们很有必要掌握它。CAS操作，当时还没有抽象出atomic包
//func cas(val *int32, old, new int32) bool

/*
	mutex的历史
	1.

*/
