package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	predic, err := prediction("sdsd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(predic)
}

func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "пн":
		return "Хуйня-1", nil
	case "вт":
		return "Хуйня-2", nil
	case "ср":
		return "Хуйня-3", nil
	case "чт":
		return "Хуйня-4", nil
	case "пт":
		return "Заебись", nil
	case "сб":
		return "Ахуенно", nil
	case "вс":
		return "Класс", nil
	default:
		return "Ты по мойму перепутал", errors.New("invalid day of week")
	}
}
