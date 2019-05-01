package main

import (
	c "common"
	"fmt"
	"math"
	"strconv"
	"time"
)

// if测试
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, lim)
	}
	return lim
}

// time
func now() string {
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		c.Pn("Today.")
	case today + 1:
		c.Pn("Tomorrow.")
	case today + 2:
		c.Pn("In two days.")
	default:
		c.Pn("Too far away.")
	}
	now := time.Now()
	switch {
	case now.Hour() < 12:
		c.Pn("Good monring!")
	case now.Hour() < 17:
		c.Pn("Good afternoon!")
	default:
		c.Pn("Good evening!")
	}
	nowStr := now.Format("2006-01-02 15:04:05")
	return nowStr
}

// defer
func myDefer() {
	defer c.Pn("The end.")
	c.Pn("Hello world!")
	for i := 0; i < 10; i++ {
		defer c.Pn(strconv.Itoa(i))
	}
}

func main() {
	myDefer()
}
