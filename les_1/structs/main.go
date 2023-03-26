package main

import "fmt"

type Age int

func (a Age) isAdult() bool {
	return a >= 18
}

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

func (u *User) printUserInfo(name string) { //(u User) - ресивер - приемнки в рамках этого метода объект скопировался
	u.name = name //но если сделать (u *User) предавть именно указатель на область памяти где хранится объект, то привоенное новое значение отразится на элементах если их использовтаь вне функции
	fmt.Println(u.name, u.age, u.sex, u.weight, u.height)
}

func (u *User) setName(name string) {
	u.name = name
}

func (u User) getName() string {
	return u.name
}

func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    Age(age), //приведение типов
		weight: weight,
		height: height,
	}
}

type DumbDatabase struct {
	m map[string]string
}

func NewDumbDatabase() *DumbDatabase {
	return &DumbDatabase{
		m: make(map[string]string),
	}
}

func main() {

	user_slice := []User{}
	user1 := NewUser("Vasya", "male", 23, 75, 185)
	user2 := NewUser("t", "sdds", 16, 75, 185)

	user_slice = append(user_slice, user1)
	user_slice = append(user_slice, user2)
	fmt.Printf("%+v\n", user1)
	fmt.Printf("%+v\n", user2)
	fmt.Println(user_slice)
	fmt.Println(user1.name)

	user1.printUserInfo("Паша") //вернула копию с перезаписанным name
	user2.printUserInfo("Паша") //вернула копию с перезаписанным name
	fmt.Println(user1)          //никаких изменений в оригинале
	fmt.Println(user2)          //никаких изменений в оригинале

	user1.setName("вадик")
	fmt.Println(user1.getName())
	user2.setName("владик")
	fmt.Println(user2.getName())

	if user1.age.isAdult() {
		fmt.Println("все можно")
	} else {
		fmt.Println("максимум пиво")
	}

	if user2.age.isAdult() {
		fmt.Println("все можно")
	} else {
		fmt.Println("максимум пиво")
	}
}
