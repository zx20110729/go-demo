package main

import (
	"fmt"
)

func testGoroute(a int, b int) int {
	sum := a + b
	fmt.Println(sum)
	return sum
}

func testPrint(a int) {
	fmt.Println(a)
}

//测试管道
func testPipe() {
	pipe := make(chan int, 3)
	pipe <- 1
	pipe <- 2
	pipe <- 3

	var num int
	num = <-pipe

	pipe <- 4

	fmt.Println("管道第一个值：", num)
	num = <-pipe
	fmt.Println("管道第二个值：", num)
	fmt.Println("管道的长度：", len(pipe))
}
