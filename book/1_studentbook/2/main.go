package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	fmt.Println()

	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
