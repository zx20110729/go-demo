package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}
type T struct {
	S string
}
type F float64

func (t T) M() {
	fmt.Println(t.S)
}
func (f F) M() {
	fmt.Println(f)
}

// 接口测试
func printInterface() {
	var i I
	i = &T{"Hello"}
	fmt.Printf("%v %T \n", i, i)
	i.M()

	i = F(math.Pi)
	fmt.Printf("%v %T \n", i, i)
	i.M()
}

//Stringer接口测试 类似toString()
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func printString() {
	p := Person{
		"zhou",
		22,
	}
	fmt.Println(p)
}
func main() {
	printString()
}
