package main

import "fmt"

func main() {
	printDefer()
}

// 1.函数
func f1(a int, b int) {
	fmt.Println(a + b)
}

//2. 参数 a b类型一致的话，可以省略前面的
func f2(a, b int) {
	fmt.Println(a + b)
}

// 3.多返回值
func swap(x, y string) (string, string) {
	return y, x
}

// 4.给返回值命名
func swap2(x, y string) (a string, b string) {
	a = y
	b = x
	return
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 5.函数作为参数传递
func printFunc() {
	var vAdd myFunc
	vAdd = add
	operator(vAdd, 1, 2)

	var vSub myFunc
	vSub = sub
	operator(vSub, 200, 100)
}

type myFunc func(int, int) int

func add(a int, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func operator(op myFunc, a, b int) {
	sum := op(a, b)
	fmt.Println(sum)
}

// 6. 可变参数 args是个slice（数组）
func sum(a int, args ...int) int {
	sum := a
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

func printSum() {
	// var nums []int = []int{1, 2, 3, 4, 5, 6}
	n := 100
	total := sum(n, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(total)
}

// defer关键字测试
func printDefer() {
	i := 0
	defer fmt.Println("defer 1 i =", i)
	i++
	defer fmt.Println("defer 2 i =", i)
	fmt.Println(i)

	myDefer(5)
}

func myDefer(a int) {
	defer fmt.Println("defer 1")
	if a == 5 {
		fmt.Println("return")
		return
	}
	defer fmt.Println("defer 2")
}
