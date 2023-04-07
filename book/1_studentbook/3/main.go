package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
}

func BenchmarkArgsString(b *testing.B) {
	for i :=0; i<
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
