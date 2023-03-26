package main

import "fmt"

func main() {
	defer handlerPanic()
	//логика defer в том чтобы откладывать выполнение функции в самый конец

	fmt.Println("main()")

	messages := []string{
		"a",
		"b",
		"c",
		"d",
	}

	messages[4] = "нахуй"
	fmt.Println(messages)
	//вызвать panic самостоятельно, при panic наше приложение падает
	// panic("нет буено")
}

func printMessage() {
	fmt.Println("printMessage()")
}

func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}

	fmt.Println("handlerPanic()")
}
