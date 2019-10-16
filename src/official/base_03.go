package official

import (
	"fmt"
	"strings"
	"math"
)

/**
更多类型：struct、slice 和映射
 */
// 1、指针
func Pointer() {
	i, j := 42, 2701
	p := &i         //指向i
	fmt.Println(*p) //通过指针读取i的值
	*p = 21         //通过指针设置i的值
	p = &j          //指向j
	*p = *p / 37    //通过指针对j进行除法运算
	fmt.Println(j)
}

// 2、结构体，就是一组字段
type Vertex struct {
	X int
	Y int
}

func PrintStruct() {
	fmt.Println(Vertex{1, 2})
}

//3、结构体字段
func PrintStruct2() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

// 4、结构体指针
func PrintStruct3() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 //使用隐式间接引用，正常使用(*p).X来访问X
	fmt.Println(v)
}

// 5、结构体文法
var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1} //Y:0 被隐式地赋予
	v3 = Vertex{}
	p  = &Vertex{1, 2} //创建一个 *Vertex 类型的结构体（指针）
)

func PrintStruct4() {
	fmt.Println(v1, p, v2, v3)
}

// 6、数组
func PrintArray() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

// 7、切片
func Slice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
}

// 8、切片就像数组的引用，更改切片会修改底层数组
func Slice2() {
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

// 9、切片文法
func Slice3() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, false, true, false}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	s2 := []Vertex{
		{1, 2},
		{3, 4},
	}
	fmt.Println(s2)
}

// 10、切片的默认行为
func Slice4() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

// 11、切片的长度与容量
func Slice5() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	//截取切片的长度使其长度为0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	//舍弃前两个值
	s = s[:2]
	printSlice(s)

}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// 12、nil切片
func Nil() {
	var s [] int
	printSlice(s)

	if s == nil {
		fmt.Println("nil!")
	}
}

// 13、用make创建切片
func Make() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

// 14、切片的切片
func Slice6() {
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	fmt.Println(board)
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	fmt.Println(board)

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// 15、向切片追加元素
func Slice7() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

// 16、Range
func Range() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//第一个参数为下标，第二个参数为对应元素的一份副本
	for i, v := range pow {
		fmt.Printf("2**%d=%d\n", i, v)
	}
}

// 17、Range 可以将下标或值赋予 _ 来忽略它。
func Range2() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Println(value)
	}
}

// 18、映射
type Vertex2 struct {
	Lat, Long float64
}

func Map() {
	var m map[string]Vertex2
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
}

// 19、映射的文法
func Map2() {
	var m = map[string]Vertex2{
		"Bell Labs": Vertex2{
			40.68433, -74.39967,
		},
		"Google": Vertex2{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

// 20、映射的文法，若顶级类型只是一个类型名
func Map3() {
	var m = map[string]Vertex2{
		"Bell Labs": {
			40.68433, -74.39967,
		},
		"Google": {
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)
}

// 21、修改映射
func Map4() {
	m := make(map[string]int)
	//赋值
	m["answer"] = 42
	fmt.Println(m["answer"])

	//通过双赋值检测某个键是否存在
	v, ok := m["answer"]
	fmt.Println(v, ok)

	//修改
	m["answer"] = 48
	fmt.Println(m["answer"])
	//删除
	delete(m, "answer")
	fmt.Println(m["answer"])

	//通过双赋值检测某个键是否存在
	v, ok = m["answer"]
	fmt.Println(v, ok)

}

// 22、函数值
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func Func() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// 23、函数的闭包
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func Func2() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg((-2 * i)))
	}
}

// 24、
// 25、
// 26、
