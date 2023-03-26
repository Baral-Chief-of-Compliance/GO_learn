package main

import (
	"errors"
	"fmt"
)

func main() {
	messages := []string{"1", "2"}
	// var notInitSlice []string
	// // notInitSlice[0] = "1" we have a panic in command line
	// fmt.Println(notInitSlice)

	notInitSlice := make([]string, 5)
	// notInitSlice := make([]string, 5, 15) //вторая цифра это copacity
	notInitSlice = append(notInitSlice, "10000") //мы расширяем слайс по
	//средством создания еще одного массива, так как слайс сделан на
	//основе массива
	//с каждым appendom вызывается преолокация массива, дорогая операция на 5 лям элм
	//
	fmt.Println(notInitSlice)
	fmt.Println(len(notInitSlice)) //длинна
	fmt.Println(cap(notInitSlice)) //емкость
	notInitSlice[3] = "100"
	fmt.Println(notInitSlice)

	printMessage(messages)
	fmt.Println(messages)
}

// в функцию передается ссылка на слайс
func printMessage(messages []string) error {
	if len(messages) == 0 {
		return errors.New("empty slice")
	}

	messages[0] = "200"
	fmt.Println(messages)
	return nil
}
