// @Author liuzhen
// @Date 2023/12/20 11:17:00
// @Desc
package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/**
goroutine

1. 进程、线程和协程
	进程(process): 程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位。
	线程(thread): 操作系统基于进程开启的轻量级进程，是操作系统调度执行的最小单位。
	协程(coroutine): 非操作系统提供而是由用户自行创建和控制的用户态‘线程’，比线程更轻量级。

2. goroutine
	Goroutine 是 Go 语言支持并发的核心，在一个Go程序中同时创建成百上千个goroutine是非常普遍的，一个goroutine会以一个很小的栈开始其生命周期，一般只需要2KB。
		区别于操作系统线程由系统内核进行调度， goroutine 是由Go运行时（runtime）负责调度。
		例如Go运行时会智能地将m个goroutine 合理地分配给n个操作系统线程，实现类似m:n的调度机制，不再需要Go开发者自行在代码层面维护一个线程池。
	Goroutine 是 Go 程序中最基本的并发执行单元。每一个 Go 程序都至少包含一个 goroutine —— main goroutine，当 Go 程序启动时它会自动创建。
	在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能——goroutine，
		当你需要让某个任务并发执行的时候，你只需要把这个任务包装成一个函数，开启一个 goroutine 去执行这个函数就可以了，就是这么简单粗暴。

	2.1. go 关键字
		一个 goroutine 必定对应一个函数/方法，可以创建多个 goroutine 去执行相同的函数/方法。
		Go语言中使用 goroutine 非常简单，只需要在函数或方法调用前加上go关键字就可以创建一个 goroutine ，从而让该函数或方法在新创建的 goroutine 中执行。
		eg: 1. go f()  // 创建一个新的 goroutine 运行函数f
			2. 匿名函数也支持使用go关键字创建 goroutine 去执行。
				go func() {
				  	// ...
				} ()

	2.2. 动态栈
		操作系统的线程一般都有固定的栈内存（通常为2MB）,而 Go 语言中的 goroutine 非常轻量级，一个 goroutine 的初始栈空间很小（一般为2KB），
		所以在 Go 语言中一次创建数万个 goroutine 也是可能的。并且 goroutine 的栈不是固定的，可以根据需要动态地增大或缩小，
		Go 的 runtime 会自动为 goroutine 分配合适的栈空间。
	2.3. goroutine调度
	2.4. GOMAXPROCS
		Go运行时的调度器使用GOMAXPROCS参数来确定需要使用多少个 OS 线程来同时执行 Go 代码。默认值是机器上的 CPU 核心数。例如在一个 8 核心的机器上，GOMAXPROCS 默认为 8。
		Go语言中可以通过runtime.GOMAXPROCS函数设置当前程序并发时占用的CPU逻辑核心数。

3. channel
	单纯地将函数并发执行是没有意义的。函数与函数间需要交换数据才能体现并发执行函数的意义。
	Go语言采用的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。
	如果说 goroutine 是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。
	Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。
		每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

	3.1. channel类型
		channel是 Go 语言中一种特有的类型。
		声明通道类型变量的格式如下: var 变量名称 chan 元素类型        其中: chan: 是关键字; 元素类型: 是指通道中传递元素的类型
		eg: var ch1 chan int   // 声明一个传递整型的通道		var ch2 chan []int // 声明一个传递int切片的通道
	3.2. channel零值: 未初始化的通道类型变量其默认零值是nil。如上为只声名但未初始化
	3.3. 初始化channel
		声明的通道类型变量需要使用内置的make函数初始化之后才能使用。具体格式如下: make(chan 元素类型, [缓冲大小])		其中: channel的缓冲大小是可选的。
		eg: ch3 := make(chan int)		ch4 := make(chan bool, 1)  // 声明一个缓冲区大小为1的通道
	3.4. channel操作: 先定义一个通道: ch := make(chan int)
		3.4.1. 发送: 将一个值发送到通道中。 ch <- 10 // 把10发送到ch中
		3.4.2. 接收: 从一个通道中接收值。 x := <- ch // 从ch中接收值并赋值给变量x			<-ch // 从ch中接收值，忽略结果
		3.4.3. 关闭: 通过调用内置的close函数来关闭通道。 close(ch)
			注意: 一个通道值是可以被垃圾回收掉的。通道通常由发送方执行关闭操作，并且只有在接收方明确等待通道关闭的信号时才需要执行关闭操作。
			它和关闭文件不一样，通常在结束操作之后关闭文件是必须要做的，但关闭通道不是必须的。
			关闭后的通道有以下特点:
				对一个关闭的通道再发送值就会导致 panic。
				对一个关闭的通道进行接收会一直获取值直到通道为空。
				对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
				关闭一个已经关闭的通道会导致 panic。

	3.5. 无缓冲的通道: 无缓冲的通道又称为阻塞的通道
		eg: 定义: ch := make(chan int)	发送: ch <- 10
		注: 创建的是无缓冲的通道，无缓冲的通道只有在有接收方能够接收值的时候才能发送成功，否则会一直处于等待发送的阶段(即一直处于阻塞状态导致程序死锁，导致panic)。
		解决: 创建一个 goroutine 去接收值:
				go func(c chan int) {
					ret := <-c
					fmt.Println("接收成功", ret)
				}
	3.6. 有缓冲的通道: 可解决上面死锁问题。在使用 make 函数初始化通道时，可以为其指定通道的容量。eg: ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	3.7. 多返回值模式: 当向通道中发送完数据时，可以通过 close() 函数来关闭通道。当一个通道被关闭后，再往该通道发送值会引发panic，从该通道取值的操作会先取完通道中的值。
		通道内的值被接收完后再对通道执行接收操作得到的值会一直都是对应元素类型的零值。多返回值用于判断一个通道是否被关闭。eg: checkChanClose
		格式: value, ok := <- ch
		其中:
			value: 从通道中取出的值，如果通道被关闭则返回对应类型的零值。
			ok: 通道打开时返回 true，通道关闭时返回false。（通道ch关闭时返回 false，否则返回 true。）

	3.8. for range接收值
		通常我们会选择使用 for range 循环从通道中接收值，当通道被关闭后，会在通道内的所有值被接收完毕后会自动退出循环。
		上面那个示例(checkChanClose)使用 for range 改写后会很简洁。
		注: 目前Go语言中并没有提供一个不对通道进行读取操作就能判断通道是否被关闭的方法。不能简单的通过len(ch)操作来判断通道是否被关闭。

	3.9. 单向通道: 通道在某个函数中只能执行发送或只能执行接收操作。
		3.9.1. 只接收通道，只能接收不能发送: <- chan int
		3.9.2. 只发送通道，只能发送不能接收: chan <- int
		注: 对一个只接收通道执行close也是不允许的，因为默认通道的关闭操作应该由发送方来完成。
	3.10. 小结: 通道操作结果。注意: 对已经关闭的通道再执行 close 也会引发 panic。
									  状态
		操作		nil			没值			有值			满
		发送		阻塞	  发送成功		   发送成功			阻塞
		接收		阻塞		阻塞		   接收成功		  接收成功
		关闭		panic	  关闭成功		   关闭成功		  关闭成功

4. select多路复用: 用于同时从多个通道接收数据，可以同时响应多个通道的操作。
	格式如下: select {
			  case <-ch1:
				  //...
			  case data := <-ch2:
				  //...
			  default:
			  	  // 默认操作
			  }
	特点:
		可处理一个或多个 channel 的发送/接收操作。
		如果多个 case 同时满足，select 会随机选择一个执行。
		对于没有 case 的 select 会一直阻塞，可用于阻塞 main 函数，防止退出。

5. 通道误用示例

6. 并发安全和锁
	6.1. 互斥锁(sync.Mutex): 互斥锁是一种常用的控制共享资源访问的方法，它能够保证同一时间只有一个 goroutine 可以访问共享资源，其他的 goroutine 则在等待锁。
		当互斥锁释放后，等待的 goroutine 才可以获取锁进入临界区，多个 goroutine 同时等待一个锁时，唤醒的策略是随机的。Go 语言中使用sync包中提供的Mutex类型来实现互斥锁。
		6.1.1. 获取互斥锁: func (m *Mutex) Lock()
		6.1.2. 释放互斥锁: func (m *Mutex) Unlock()
	6.2. 读写互斥锁(sync.RWMutex): 适用于读多写少场景。读取资源不加锁，写资源加锁。
		6.2.1. 获取写锁:func (rw *RWMutex) Lock()
		6.2.2. 释放写锁: func (rw *RWMutex) Unlock()
		6.2.3. 获取读锁: func (rw *RWMutex) RLock()
		6.2.4. 释放读锁: func (rw *RWMutex) RUnlock()
		6.2.5. 返回一个实现Locker接口的读写锁: func (rw *RWMutex) RLocker() Locker
	6.3. sync.WaitGroup: 在代码中生硬的使用time.Sleep肯定是不合适的，Go语言中可以使用sync.WaitGroup来实现并发任务的同步。
		6.3.1. 计数器+delta: func (wg * WaitGroup) Add(delta int)
		6.3.2. 计数器-1: (wg *WaitGroup) Done()
		6.3.3. 阻塞直到计数器变为0: (wg *WaitGroup) Wait()
		sync.WaitGroup内部维护着一个计数器，计数器的值可以增加和减少。例如当我们启动了 N 个并发任务时，就将计数器值增加N。
		每个任务完成时通过调用 Done 方法将计数器减1。通过调用 Wait 来等待并发任务执行完，当计数器值为 0 时，表示所有并发任务已经完成。
		注: 需要注意sync.WaitGroup是一个结构体，进行参数传递的时候要传递指针。
	6.4. sync.Once: 在某些场景下我们需要确保某些操作即使在高并发的场景下也只会被执行一次，例如只加载一次配置文件等。
		Go语言中的sync包中提供了一个针对只执行一次场景的解决方案——sync.Once，sync.Once只有一个Do方法，
		其签名如下: func (o *Once) Do(f func())
		注: 如果要执行的函数f需要传递参数就需要搭配闭包来使用。
		使用场景: 加载配置文件；并发安全的单例模式
		实现原理: 内部包含一个互斥锁和一个布尔值，互斥锁保证布尔值和数据的安全，而布尔值用来记录初始化是否完成。这样设计就能保证初始化操作的时候是并发安全的并且初始化操作也不会被执行多次。
	6.5. sync.Map: 并发安全map。箱即用表示其不用像内置的 map 一样使用 make 函数初始化就能直接使用。同时sync.Map内置了诸如Store、Load、LoadOrStore、Delete、Range等操作方法。
		6.5.1 存储key-value数据: func (m *Map) Store(key, value interface{})
		6.5.2 查询key对应的value: func (m *Map) Load(key interface{}) (value interface{}, ok bool)
		6.5.3 查询或存储key对应的value: func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
		6.5.4 查询并删除key: func (m *Map) LoadAndDelete(key interface{}) (value interface{}, loaded bool)
		6.5.5 删除key: func (m *Map) Delete(key interface{})
		6.5.6 对map中的每个key-value依次调用f: func (m *Map) Range(f func(key, value interface{}) bool)

7. 原子操作: 针对整数数据类型（int32、uint32、int64、uint64）我们还可以使用原子操作来保证并发安全，通常直接使用原子操作比使用锁操作效率更高。
		Go语言中原子操作由内置的标准库 sync/atomic 提供。
	读取操作:
		func LoadInt32(addr *int32) (val int32)
		func LoadInt64(addr *int64) (val int64)
		func LoadUint32(addr *uint32) (val uint32)
		func LoadUint64(addr *uint64) (val uint64)
		func LoadUintptr(addr *uintptr) (val uintptr)
		func LoadPointer(addr *unsafe.Pointer) (val unsafe.Pointer)
	写入操作:
		func StoreInt32(addr *int32, val int32)
		func StoreInt64(addr *int64, val int64)
		func StoreUint32(addr *uint32, val uint32)
		func StoreUint64(addr *uint64, val uint64)
		func StoreUintptr(addr *uintptr, val uintptr)
		func StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)
	修改操作:
		func AddInt32(addr *int32, delta int32) (new int32)
		func AddInt64(addr *int64, delta int64) (new int64)
		func AddUint32(addr *uint32, delta uint32) (new uint32)
		func AddUint64(addr *uint64, delta uint64) (new uint64)
		func AddUintptr(addr *uintptr, delta uintptr) (new uintptr)
	交换操作:
		func SwapInt32(addr *int32, new int32) (old int32)
		func SwapInt64(addr *int64, new int64) (old int64)
		func SwapUint32(addr *uint32, new uint32) (old uint32)
		func SwapUint64(addr *uint64, new uint64) (old uint64)
		func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)
		func SwapPointer(addr *unsafe.Pointer, new unsafe.Pointer) (old unsafe.Pointer)
	比较并交换操作:
		func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
		func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)
		func CompareAndSwapUint32(addr *uint32, old, new uint32) (swapped bool)
		func CompareAndSwapUint64(addr *uint64, old, new uint64) (swapped bool)
		func CompareAndSwapUintptr(addr *uintptr, old, new uintptr) (swapped bool)
		func CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) (swapped bool)
*/

var (
	x int64
	// 等待组
	wg sync.WaitGroup
	// 互斥锁
	m sync.Mutex

	mutex   sync.Mutex
	rwMutex sync.RWMutex
)

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	// 3.7 判断一个通道是否被关闭
	checkChanClose(ch)
	// 3.8
	forRangeDemo(ch)

	// 3.9.
	ch2 := Producer2()
	res2 := Consumer2(ch2)
	fmt.Println(res2) // 25

	// 6.1. 互斥锁(sync.Mutex)
	wg.Add(2)
	go addAndLock()
	go addAndLock()
	wg.Wait()
	fmt.Println(x)

	// 6.2. 读写互斥锁(sync.RWMutex)
	// 使用互斥锁，10并发写，1000并发读
	do(writeWithLock, readWithLock, 10, 1000) // x:10 cost:1.466500951s
	// 使用读写互斥锁，10并发写，1000并发读
	do(writeWithRWLock, readWithRWLock, 10, 1000) // x:10 cost:117.207592ms
}

// 3.7. 多返回值模式， 判断一个通道是否被关闭
func checkChanClose(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

// 3.8. for range接收值，取代 checkChanClose
func forRangeDemo(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

// 3.9.1. Producer2 返回一个接收通道
func Producer2() <-chan int {
	ch := make(chan int, 2)
	// 创建一个新的goroutine执行发送数据的任务
	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 1 {
				ch <- i
			}
		}
		close(ch) // 任务完成后关闭通道
	}()

	return ch
}

// 3.9.2. Consumer2 参数为接收通道
func Consumer2(ch <-chan int) int {
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

// 4. select 实例
func selectDemo() {
	// 这种方式虽然可以实现从多个通道接收值的需求，但是程序的运行性能会差很多。
	/*ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	for {
		// 尝试从ch1接收值
		data, ok := <-ch1
		// 尝试从ch2接收值
		data, ok := <-ch2
	}*/

	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		// 接收
		case x := <-ch:
			fmt.Println(x)
		// 发送
		case ch <- i:
		}
	}
}

// 6.1. 互斥锁(sync.Mutex) addAndLock 对全局变量x执行5000次加1操作
func addAndLock() {
	for i := 0; i < 5000; i++ {
		m.Lock() // 修改x前加锁
		x = x + 1
		m.Unlock() // 改完解锁
	}
	wg.Done()
}

// 6.2. 读写互斥锁(sync.RWMutex)
// writeWithLock 使用互斥锁的写操作
func writeWithLock() {
	mutex.Lock() // 加互斥锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	mutex.Unlock()                    // 解互斥锁
	wg.Done()
}

// readWithLock 使用互斥锁的读操作
func readWithLock() {
	mutex.Lock()                 // 加互斥锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	mutex.Unlock()               // 释放互斥锁
	wg.Done()
}

// writeWithLock 使用读写互斥锁的写操作
func writeWithRWLock() {
	rwMutex.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwMutex.Unlock()                  // 释放写锁
	wg.Done()
}

// readWithRWLock 使用读写互斥锁的读操作
func readWithRWLock() {
	rwMutex.RLock()              // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwMutex.RUnlock()            // 释放读锁
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	// wc个并发写操作
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	//  rc个并发读操作
	for i := 0; i < rc; i++ {
		wg.Add(1)
		go rf()
	}

	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x:%v cost:%v\n", x, cost)

}

// 6.4. sync.Once
type singleton struct{}

var instance *singleton
var once sync.Once

// 并发安全的单例模式
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

// 6.5. 并发安全map
func syncMapDemo() {
	// 并发安全的map
	var m = sync.Map{}

	wg := sync.WaitGroup{}
	// 对m执行20个并发的读写操作
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)         // 存储key-value
			value, _ := m.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()

}
