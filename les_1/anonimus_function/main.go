package main

import (
	"fmt"
)

func main() {

	func() {
		fmt.Println("анонимная функция")
	}()
	//замыкание - функция, которая ссылается на незаивисмые свободные переменные
	//при замыкание функция запоминает состояние при котором она была создана

	inc := incremant()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	inc = incremant()
	fmt.Println(inc())
	fmt.Println(inc())

}

func incremant() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
