package main

import (
	"errors"
	"fmt"
	"reflect"
)

var b, c, d int

func main() {

	message := []byte("asd")
	var a rune = 'a'
	fmt.Println(message)
	fmt.Println(reflect.TypeOf(message))

	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	b, c, d = 1, 2, 3

	b, c = c, b
	fmt.Println(b, c, d)

	for i := 17; i < 20; i += 1 {
		result := fmt.Sprintf("%d", i)
		print(result)
		fmt.Println(sayHello("иди нахуй"))
		forFunc, err := checkAge(i)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(forFunc)

	}

}

func print(message string) {
	fmt.Println(message)
}

func sayHello(name string) string {
	result := fmt.Sprintf("Привет, %s", name)
	return result
}

func checkAge(age int) (string, error) {
	if age >= 18 && age < 45 {
		response := fmt.Sprintf("Тебе %d все можно", age)
		return response, nil
	} else if age >= 45 {
		response := fmt.Sprintf("Тебе %d так что только спать", age)
		return response, errors.New("yo are too old")
	}

	response := fmt.Sprintf("Тебе %d так что только пиво", age)
	return response, errors.New("you are too young")
}
