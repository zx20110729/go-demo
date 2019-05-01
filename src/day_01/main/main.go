package main

import (
	"day_01/cacl"
	"day_01/goroute"
	"fmt"
	"time"
)

func main() {
	sum := cacl.Add(1, 2)
	fmt.Println("sum =", sum)
	sub := cacl.Sub(3, 1)
	fmt.Println("sub =", sub)

	pipe := make(chan int, 1)
	go goroute.Add(1, 6, pipe)

	goSum := <-pipe
	fmt.Println("goSum =", goSum)
	time.Sleep(time.Second)
}
