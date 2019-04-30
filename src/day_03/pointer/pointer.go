package main

import "fmt"

//指针类型测试
func printPointer() {
	var a int = 5
	var p *int = &a
	fmt.Println(p)
	fmt.Println(*p)
	//a也会被修改
	*p = 100
	fmt.Println(a)
}
func main() {
	printPointer()
}
