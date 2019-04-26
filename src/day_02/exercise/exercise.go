package main

import (
	"fmt"
	r "math/rand"
	"os"
	"time"
)

const (
	male   = "male"
	female = "female"
)

// 打印 加法
func printAdd(sum int) {
	for i := 0; i <= sum; i++ {
		other := sum - i
		fmt.Println(i, "+", other, "=", sum)
	}
}

// 根据当前时间打印不同的信息
func printGender() {

	for count := 1; count < 20; count++ {
		now := time.Now().Unix()
		var gender string
		if now%2 == 0 {
			gender = male
		} else {
			gender = female
		}
		fmt.Println("gender =", gender)
		time.Sleep(time.Second)
	}
}

//打印系统信息
func printOs() {
	var goPath string = os.Getenv("GOPATH")
	path := os.Getenv("PATH")
	fmt.Println("goPath =", goPath)
	fmt.Println("path=", path)
}

// 值类型和引用类型
func printReference() {
	a := 1
	b := 1
	fmt.Println("before modified a =", a)
	fmt.Println("before modified b =", b)
	modify(a)
	// &b：获取b的地址
	modifyPointer(&b)
	fmt.Println("after modified a =", a)
	fmt.Println("after modified b =", b)
}
func modify(a int) {
	a = a + 1
}
func modifyPointer(a *int) {
	// *a：获取a指向的地址存储的值
	*a = *a + 1
}

// 交换两个int值
func swapInt(a *int, b *int) {
	tmp := *b
	*b = *a
	*a = tmp
}

// 类型转换 var a int = 1; var b int8 = int8(a)
func typeConvert() {
	var a int8 = 127
	var b int16 = int16(a)
	fmt.Printf("a=%d,b=%d\n", a, b)

	var c int = 1
	var d int32 = int32(c + c)
	fmt.Printf("c=%d,d=%d", c, d)

}

//生成10个随机数
func rand() {
	r.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		var a int = r.Intn(100)
		fmt.Print(a)
		if i == 9 {
			fmt.Println()
		} else {
			fmt.Printf(",")
		}
	}

	for j := 0; j < 10; j++ {
		var a float32 = r.Float32()
		fmt.Print(a)
		if j == 9 {
			fmt.Println()
		} else {
			fmt.Printf(",")
		}
	}
}

// 打印字符串类型
func printString() {
	s := `abc\n`
	s2 := "abc\n"
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println("--------")
	// %q打印带类型  "hello"
	str := "hello"
	fmt.Printf("%q\n", str)
}

//打印byte类型
func printByte() {
	var c byte = 'c'
	fmt.Println("c =", c)
	fmt.Printf("c = %c\n", c)
}

func main() {
	printByte()

}
