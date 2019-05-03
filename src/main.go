package main

import (
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
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	now := time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Println("Good monring!")
	case now.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
	nowStr := now.Format("2006-01-02 15:04:05")
	return nowStr
}

// defer
func myDefer() {
	defer fmt.Println("The end.")
	fmt.Println("Hello world!")
	for i := 0; i < 10; i++ {
		defer fmt.Println(strconv.Itoa(i))
	}
}

//指针
func pointer() {
	i, j := 42, 2701
	p := &i
	fmt.Println(strconv.Itoa(*p))
	*p = 21
	fmt.Println(strconv.Itoa(i))

	p = &j
	*p = *p / 37
	fmt.Println(strconv.Itoa(j))
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
	fmt.Println(v)
	//隐式间接调用
	p := &v
	p.X = 100
	fmt.Println(v.X)
	fmt.Println(p.Y)
	//打印成员变量
	fmt.Println("v1 =", v1)
	fmt.Println("v2 =", v2)
	fmt.Println("v3 =", v3)
	fmt.Println("p1 =", *p1)
}

//数据和分片
func arrayAndSlice() {
	//数组
	var arr [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
	s := arr[1:8]
	s = append(s, 10, 11, 12)
	fmt.Println(s)
	fmt.Printf("%T,%T\n", arr, s)
	a := make([]int, 5)
	b := make([]int, 0, 5)
	fmt.Println(len(a), cap(a))
	fmt.Println(len(b), cap(b))
}

//range
func myRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Println("i =", i, " v =", v)
	}
}

type User struct {
	Name string
	Age  int
}

//map测试
func myMap() {
	//声明1
	var m map[string]User
	m = make(map[string]User)
	m["1"] = User{"zhou", 1}
	fmt.Println(m["1"].Name)
	//声明2
	var m2 = map[string]User{
		"1": User{
			Name: "zhou2",
			Age:  2,
		},
		"2": User{
			"pan2", 2,
		},
	}
	fmt.Println(m2["2"].Name)
	//声明3
	var m3 = map[string]User{
		"1": {"zhou3", 3},
		"2": {"pan3", 3},
	}
	fmt.Println(m3["1"].Name)
}

func main() {
	myMap()
}
