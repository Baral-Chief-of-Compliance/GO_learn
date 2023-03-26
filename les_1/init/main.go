package main

import "fmt"

var msg string

// всегда срабатывает перед вызовом функции main
func init() {
	msg = "здарова бандиты"
}

func main() {
	fmt.Println(msg)
}
