package main

import (
	"fmt"
	"io"
	"strings"
)

func printReader() {
	r := strings.NewReader("Hello world!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, b = %v, err = %v.\n", n, b, err)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}
}
func main() {
	printReader()

}
