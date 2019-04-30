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
func main() {
	printAtoi()
}
