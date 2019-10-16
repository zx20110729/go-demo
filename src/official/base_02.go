package official

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

/**
流程控制语句：for、if、else、switch 和 defer
 */
// 1、for
func PrintFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 2、for，初始化语句和后置语句是可选的。
func PrintFor2() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

// 3、for是Go中的"while"
func PrintFor3() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// 4、无限循环
func PrintFor4() {
	for {
		fmt.Println(1)
	}
}

// 5、if
func Sqrt(x float64) string {
	if x < 0 {
		return Sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// 6、if的简短语句，if 语句可以在条件表达式前执行一个简单的语句
func Pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// 7、if和else
func Pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
		return lim
	}
}

// 8、switch
func PrintSwitch() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s,\n", os)
	}
}

// 9、switch求值顺序
func PrintSwitch2() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Toady.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// 10、没有条件的switch,没有条件的 switch 同 switch true 一样
func PrintSwitch3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

// 11、defer,defer 语句会将函数推迟到外层函数返回之后执行。
func PrintDefer() {
	defer fmt.Println("world")
	fmt.Println("Hello")
}

// 12、defer栈
func PrintDefer2() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
