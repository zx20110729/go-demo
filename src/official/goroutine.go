package official

import (
	"fmt"
	"time"
	"sync"
)

/**
Go 程（goroutine）是由 Go 运行时管理的轻量级线程。
 */

// 1、Go程
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}
func Go() {
	go say("world")
	say("Hello")
}

// 2、信道 信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
func sum(n []int, c chan int) {
	sum := 0
	for _, v := range n {
		sum += v
	}
	c <- sum
}
func Go2() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //从c中接收
	fmt.Println(x, y, x+y)
}

//3、带缓冲的信道，将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：
func Go3() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// 4、信道的range和close
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func Go4() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

// 5、select语句，语句使一个 Go 程可以等待多个通信操作
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y := y, x+y
			fmt.Sprintf("%v,%v", x, y)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func Select() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci2(c, quit)
}

// 6、默认选择，当 select 中的其它分支都没有准备好时，default 分支就会执行。
func Select2() {
	tick := time.Tick(100 * time.Microsecond)
	boom := time.After(500 * time.Microsecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("boom.")
		default:
			fmt.Println("default.")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// 7、sync.Mutex 同步锁

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func Go5() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
