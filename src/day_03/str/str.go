package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 比较string类型前缀是否一致
func printHasPrefix(str string, prefix string) bool {
	var result bool = strings.HasPrefix(str, prefix)
	return result
}

// 比较string类型后缀是否一致
func printHasSuffix(str string, suffix string) bool {
	var result bool = strings.HasSuffix(str, suffix)
	return result
}

// 去除首尾空格
func printTrim() {
	var spaceStr string = "  hello  "
	result := strings.TrimSpace(spaceStr)
	fmt.Println("origin string:", spaceStr, ",after trim:", result, ".")
}

//分割字符串为数组
func printSplit() {
	str := "a,b,c,d,e,f"
	var strArr []string
	strArr = strings.Split(str, ",")
	fmt.Println(strArr)
	fmt.Println(len(strArr))
}

// 连接数组间的元素
func printJoin() {
	var strArr []string = []string{"a", "b", "c", "d"}
	var sep string = "-"
	result := strings.Join(strArr, sep)
	fmt.Println(result)
}

//把一个整数转换为字符串
func printItoa() {
	var num int = 100
	var numStr string = strconv.Itoa(num)
	fmt.Println(numStr)
	fmt.Printf("%q\n", numStr)
}

//字符串转数字
func printAtoi() {
	var numStr string = "100"
	num, _ := strconv.Atoi(numStr)
	fmt.Println(num)
	fmt.Printf("%q\n", num)
}

// string 类型测试
func printString() {
	str := "http://www.baidu.com"
	prefix := "http://"
	suffix := ".com"
	fmt.Println(str, "is start with", prefix, ":", printHasPrefix(str, prefix))
	fmt.Println(str, "is end with", suffix, ":", printHasSuffix(str, suffix))
}

// 回文判断
func process(str string) bool {
	t := []rune(str)
	length := len(t)
	if length <= 1 {
		return true
	}
	for i := 0; i < length-1; i++ {
		if t[i] != t[length-1] {
			return false
		}
		length--
	}
	return true
}

//统计字符串里的字母 数字 和空格的个数
func count(str string) (wordCount, spaceCount, otherCount int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			wordCount++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			otherCount++
		}
	}
	fmt.Printf("wordCount: %d, spaceCount: %d, otherCount: %d.\n", wordCount, spaceCount, otherCount)
	return
}
func main() {
	// var str string = "Hello world! 123456"
	var str string = "上海自来a水aa来自海上"
	result := process(str)
	fmt.Println("result:", result)
}
