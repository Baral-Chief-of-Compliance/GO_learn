package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMin(1, 2, 3, 55555, 333))
	fmt.Println(findMax(1, 2, 3, 55555, 333))
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}

	return min
}

func findMax(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]

	for _, i := range numbers {
		if i > max {
			max = i
		}
	}

	return max
}
