package main

import "fmt"

func main() {

	number := 5
	var p *int
	fmt.Println(p)

	p = &number

	fmt.Println(p)
	fmt.Println(&number)

	message := "я клоун"
	fmt.Println(message)
	printMessage(&message)
	fmt.Println(message)

}

func printMessage(message *string) {
	*message += "(из функции printMessage)"
	fmt.Println(*message)
}
