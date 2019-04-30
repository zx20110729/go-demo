package main

import (
	"fmt"
	"time"
)

// 时间单位
func unit() {
}

// 格式化
func format() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}

//执行耗时
func duration() {
	start := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
	}
	end := time.Now().UnixNano()
	fmt.Println("for spend", (end-start)/1000/1000, "ms.")
}
func main() {
	duration()
}
