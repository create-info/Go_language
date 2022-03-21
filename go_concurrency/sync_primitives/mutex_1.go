package sync_primitives

import (
	"fmt"
	"sync"
)

/*
SumWithOutMutex synchronization primitives mutex
*/
func SumWithOutMutex() {
	var count = 0
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)

}

func SumWithMutes() {
	// 互斥锁保护计数器
	var counter Counter

	// 辅助变量，用来确认所有的goroutine都完成
	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个go routine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加10万次
			for j := 0; j < 100000; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}

//Counter “多个goroutine并发更新同一个资源”，这里的同一个资源的条件应该是堆栈分析后分配到堆上的变量，因为堆上是线程共享的，如果是栈上的变量的话，因为是线程独有的就不会出现并发更新的问题.
type Counter struct {
	CounterType int
	Name        string
	mu          sync.Mutex
	count       uint64
}

//Incr 加1的方法，内部使用互斥锁保护
func (c *Counter) Incr() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

//Count 得到计数器的值，也需要锁保护, 思考为什么读也要加锁？
/*
1.mutex保护的临界区。如果读的时候不加锁，可能会造成不一致的后果，比如部分变量被修改了。
2.如果临界区比较简单，比如一个int64读写，也可能在一些cpu架构下有可见性问题，导致别的goroutine对变量的写读goroutine看不到
 在一个cpu上的写不一定及时的同步到另一个cpu核.amd64的可能没问题，arm就不一定了。看go内存模型
*/
func (c *Counter) Count() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

/*
思考：如果 Mutex 已经被一个 goroutine 获取了锁，其它等待中的 goroutine 们只能一直等待。那么，等这个锁释放后，等待中的 goroutine 中哪一个会优先获取 Mutex 呢？
等待的goroutine们是以FIFO排队的 (from https://go.dev/src/sync/mutex.go)
1）当Mutex处于正常模式时，若此时没有新goroutine与队头goroutine竞争，则队头goroutine获得。若有新goroutine竞争大概率新goroutine获得。
2）当队头goroutine竞争锁失败1ms后，它会将Mutex调整为饥饿模式。进入饥饿模式后，锁的所有权会直接从解锁goroutine移交给队头goroutine，此时新来的goroutine直接放入队尾。
3）当一个goroutine获取锁后，如果发现自己满足下列条件中的任何一个
	#1 它是队列中最后一个
	#2 它等待锁的时间少于1ms，则将锁切换回正常模式

鸟叔的：sync.mutex 源代码分析
https://colobu.com/2018/12/18/dive-into-sync-mutex/

golang源码阅读-sync.Mutex
https://studygolang.com/articles/17017

go 语言查看汇编代码命令:
go tool compile -S mutex_1.go
对于counter 的例子可以优化一下，那就是获取计数的 Count 函数其实可以通过读写锁，也就是 RWMutex 来优化一下。

运行代码是加上 -race 参数可以使用相关工具检测资源争抢的问题。go run -race xx.go  运行时检测是否有数据竞争问题

*/
