package main

import (
	"errors"
	"fmt"
	"time"
)

// panic抛出异常
func printPanic() {
	err := errors.New("init config error")
	if err != nil {
		panic(err)
	}
	return
}

//数组拼接函数
func printAppend() {
	var arr []int
	arr = append(arr, 1, 2, 3, 4, 5, 6)
	fmt.Println(arr)
}

// 异常处理函数
func printRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	b := 0
	a := 100 / b
	fmt.Println(a)
	return
}

// new和make函数测试
func printNewAndMake() {
	s1 := new([]int)
	*s1 = make([]int, 5)
	(*s1)[0] = 100
	fmt.Println(s1)
	s2 := make([]int, 10)
	s2[0] = 200
	fmt.Println(s2)
}

//递归测试
func recursive(n int) {
	if n > 10 {
		return
	}
	fmt.Println(n, "- Hello")
	time.Sleep(time.Second)
	recursive(n + 1)
}

//斐波那契数
func fab(n int) int {
	if n <= 1 {
		return 1
	}
	sum := fab(n-1) + fab(n-2)
	return sum
}

//闭包
func closePackage() {
	f := adder()
	fmt.Println(f(1))
	fmt.Println(f(100))
	fmt.Println(f(1000))
}

func adder() func(int) int {
	//x会被返回的函数引用共享
	var x int
	return func(num int) int {
		x += num
		return x
	}
}
func main() {
	closePackage()
}
