package main

import "fmt"

func main() {
	messages := []string{
		"a",
		"b",
		"c",
		"d",
		"f",
	}

	for i := 0; i < len(messages); i++ {
		fmt.Println(i)
	}

	for i := range messages {
		fmt.Println(i)
	}

	for _, value := range messages {
		fmt.Println(value)
	}

	counter := 0

	for {
		if counter == 100 {
			break
		}
		counter++
		fmt.Println(counter)
	}
}
