package shape

import (
	"fmt"
	"math"
)

type Shape interface {
	ShapeWithArea
	ShapeWithPerimetr
}

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimetr interface {
	Perimetr() float32
}

type Square struct {
	sideLenght float32
}

func NewSquare(lenght float32) Square { //конструктор
	return Square{
		sideLenght: lenght,
	}
}

func (s Square) Area() float32 {
	return s.sideLenght * s.sideLenght
}

func (s Square) Perimetr() float32 {
	return 4 * s.sideLenght
}

type Circle struct {
	radius float32
}

func NewCircle(radius float32) Circle {
	return Circle{
		radius: radius,
	}
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) Perimetr() float32 {
	return 2 * math.Pi * c.radius
}

func PrintShapeArea(shape Shape) {
	fmt.Println("Площадь: ", shape.Area())
	fmt.Println("Периметр: ", shape.Perimetr())
}

func PrintInterface(i interface{}) {
	switch value := i.(type) { //приведение интерфейса к конкретному типу
	case int:
		fmt.Println("int", value)
	case bool:
		fmt.Println("bool", value)
	default:
		fmt.Println("unknown type", value)
	}

	str, ok := i.(string) //приведение интерфейса к конкретному типу
	if !ok {
		fmt.Println("interface not string")
	}
	fmt.Println(str)
	fmt.Println(len(str))
	fmt.Printf("%+v\n", i)
}
