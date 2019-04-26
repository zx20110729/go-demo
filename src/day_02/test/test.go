package test

import (
	_ "day_02/test2"
	"fmt"
)

var Name string = "test"
var Age int

func init() {
	Name = "Hello  world!"
	Age = 10
	fmt.Println("test init...")
}
