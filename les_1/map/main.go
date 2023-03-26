package main

import "fmt"

func main() {
	users := map[string]int{
		"bob": 15,
		"gob": 48,
	}

	age, exists := users["bob"]
	if exists {
		fmt.Println("bob", age)
	} else {
		fmt.Println("bob нет в списке")
	}

	fmt.Println(users["bob"])

	users["Serega"] = 21
	for key, value := range users {
		fmt.Println(key, value)
	}

	delete(users, "bob")

	for key, value := range users {
		fmt.Println(key, value)
	}

	var usersD map[string]int
	//мы можем использовать тот же make для инциализации map
	usersD = make(map[string]int)
	fmt.Println(usersD)
	usersD["vasya"] = 15
	fmt.Println(usersD)
	//если map не инициализировна, мы в нее не можем вставлять пары ключ-значение

	//к map можно применять встроенную функцию len, a cap нельзя, так как он тольок для slice
	fmt.Println(len(users))
}
