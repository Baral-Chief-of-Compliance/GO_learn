package main

import (
	"fmt"
	"os"
	"testing"
)

func main() {
	var s, sep string
	// fmt.Println(s, sep)
	// fmt.Println(len(os.Args))
	// for i := 0; i < len(os.Args); i++ {
	// 	fmt.Println(os.Args[i])
	// }
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func BenchmarkMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
