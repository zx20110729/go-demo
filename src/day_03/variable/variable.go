package main

import "fmt"

//变量声明
var c, python, java bool

//变量初始化
var a, b int = 1, 2

func variable() {
	//短变量声明
	//函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。
	i := 100
	fmt.Println(i, c, python, java)
	fmt.Println(a, b)
}
func main() {
	variable()
}
