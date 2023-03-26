package main

import (
	"errors"
	"fmt"
)

func main() {
	messages := [3]string{"1", "2", "3"}
	fmt.Println(messages)
	fmt.Println(messages[1])
	fmt.Println(messages[1])
	printMessage(messages)
	fmt.Println(messages)
}

// в функцию передается копия массива
func printMessage(messages [3]string) error {
	if len(messages) == 0 {
		return errors.New("empty array")
	}

	messages[0] = "200"

	fmt.Println(messages)

	return nil
}
