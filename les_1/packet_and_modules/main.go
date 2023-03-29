package main

import (
	"les_2/shape"
)

func main() {
	square := shape.NewSquare(5)
	circle := shape.NewCircle(8)

	// scheduler := scheduler.NewScheduler()

	shape.PrintShapeArea(square)
	shape.PrintShapeArea(circle)

	shape.PrintInterface(square)
	shape.PrintInterface(circle)
	shape.PrintInterface(true)
	shape.PrintInterface(2)
	shape.PrintInterface("sss")
}
