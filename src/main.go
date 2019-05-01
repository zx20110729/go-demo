package main

import (
	c "common"
	"fmt"
	"math"
	"runtime"
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

func main() {
	c.Pn(runtime.GOOS)
}
