package official

import (
	"fmt"
	"math/rand"
	"math"
	"math/cmplx"
)

/**
包、变量和函数
 */
// 1
func HelloWorld() {
	fmt.Println("Hello World.")
}

// 2、包
func PrintRand() {
	fmt.Println("My favorite number is ", rand.Intn(10))
}

// 3、导入
func PrintSqrt() {
	fmt.Println("My favorite number is ", math.Sqrt(9))
}

// 4、导出名，首字母大写的名字表示是导出的
func PrintPi() {
	fmt.Println(math.Pi)
}

// 5、函数
func Add(x int, y int) int {
	return x + y
}

// 6、函数，参数类型省略
func Add2(x, y int) int {
	return x + y
}

// 7、多值返回
func Swap(x, y string) (string, string) {
	return y, x
}

// 8、命名返回值
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 9、变量
func PrintVar() {
	var c, python, java bool
	var i int
	fmt.Println(i, c, python, java)
}

// 10、变量的初始化
func PrintVar2() {
	var i, j int = 1, 2
	var c, python, java = true, true, "no!"
	fmt.Println(i, j, c, python, java)
}

// 11、短变量声明
//函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用
func PrintVar3() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}

// 12、基本类型
// bool、string、int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64、uintptr、byte(uint8的别名)、rune(int32的别名)、float32、float64、complex64、complex128
var (
	Tobe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func PrintVar4() {
	fmt.Printf("Type: %T Value: %v\n", Tobe, Tobe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// 13 、零值
func PrintZero() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println(i, f, b, s)
}

// 14、类型转换
func PrintVar5() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt((float64((x*x + y*y))))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// 15、类型推导
func PrintVar6() {
	v := 42
	fmt.Printf("v is of type %T", v)
}

// 16、常量
const Pi = 3.14

func PrintConst() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

// 17、数值常量
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func PrintConst2() {
	fmt.Println(Small*10 + 1)
	fmt.Println(Small * 0.1)
	fmt.Println(Big * 0.1)
	//fmt.Println(Big*10 + 1) //此行出现异常，int值越界，Big常量通过上下文推导出为int
}
