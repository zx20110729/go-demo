package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	printLable()
}

//从控制台读取一个数值并打印
func readConsole() {
	var s string
	fmt.Scanf("%s", &s)

	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%T %d \n", num, num)
	}
	rand.Seed(time.Now().Unix())
	numRand := rand.Intn(100)
	switch numRand {
	case num:
		fmt.Println("中奖了")
	default:
		fmt.Println("中奖号码为：", numRand)
	}
}

// switch
func printSwitch() {
	a := 0
	switch a {
	case 0:
		fmt.Println(0)
		fmt.Println(1)
		fmt.Println(9)
		//走到下一个分支
		fallthrough
	case 10:
		fmt.Println(10)
	default:
		fmt.Println("default")
	}
}

func printFor() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}
}

func printForRange() {
	str := "Hello world! From 中国😄😳"
	for i, v := range str {
		fmt.Printf("index: %d , value: %c , len: %d .\n", i, v, len([]byte(string(v))))
	}
}

func printLable() {
outer:
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			if j == 4 {
				fmt.Println()
				continue outer
			}
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}
}
