package main

import (
	"fmt"
	"time"
)

func add(a int, b int) int {
	var sum int
	sum = a + b
	return sum
}

func main() {
	var sum int
	sum = add(100, 5)
	fmt.Println("sum = ", sum)
	//goroute测试
	// go testGoroute(100, 200)

	// for i := 0; i < 100; i++ {
	// 	go testPrint(i)
	// }
	//主线程等待子线程执行
	// time.Sleep(time.Second)

	go testPipe()

	time.Sleep(time.Second)
}
