package common

import "fmt"

//打印并换行
func Pn(a ...interface{}) (n int, err error) {
	return fmt.Println(a)
}
