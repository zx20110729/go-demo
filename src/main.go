package main

import (
	c "common"
	"fmt"
	"math"
	"strconv"
	"time"
)

// if测试
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, lim)
	}
	return lim
}

// time
func now() string {
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		c.Pn("Today.")
	case today + 1:
		c.Pn("Tomorrow.")
	case today + 2:
		c.Pn("In two days.")
	default:
		c.Pn("Too far away.")
	}
	now := time.Now()
	switch {
	case now.Hour() < 12:
		c.Pn("Good monring!")
	case now.Hour() < 17:
		c.Pn("Good afternoon!")
	default:
		c.Pn("Good evening!")
	}
	nowStr := now.Format("2006-01-02 15:04:05")
	return nowStr
}

// defer
func myDefer() {
	defer c.Pn("The end.")
	c.Pn("Hello world!")
	for i := 0; i < 10; i++ {
		defer c.Pn(strconv.Itoa(i))
	}
}

//指针
func pointer() {
	i, j := 42, 2701
	p := &i
	c.Pn(strconv.Itoa(*p))
	*p = 21
	c.Pn(strconv.Itoa(i))

	p = &j
	*p = *p / 37
	c.Pn(strconv.Itoa(j))
}

//结构体
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  //创建一个Vertex类型结构体
	v2 = Vertex{X: 1}  //Y:0 被隐式的赋予
	v3 = Vertex{}      //X:0 Y:0
	p1 = &Vertex{1, 2} //创建一个*Vertex类型的结构体（指针）
)

func myStruct() {
	var v Vertex
	v = Vertex{1, 2}
	c.Pn(v)
	//隐式间接调用
	p := &v
	p.X = 100
	c.Pn(v.X)
	c.Pn(p.Y)
	//打印成员变量
	c.Pn("v1 =", v1)
	c.Pn("v2 =", v2)
	c.Pn("v3 =", v3)
	c.Pn("p1 =", *p1)
}

//数据和分片
func arrayAndSlice() {
	//数组
	var arr [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	c.Pn(arr)
	s := arr[1:8]
	c.Pn(s)
	fmt.Printf("%T,%T\n", arr, s)
}
func main() {
	arrayAndSlice()
}
