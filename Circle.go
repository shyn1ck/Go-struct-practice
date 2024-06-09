package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
	Color  string
}

func main() {
	circle := Circle{
		Radius: 5,
		Color:  "Красный",
	}
	area := circle.Area()
	circumference := circle.Circumference()
	fmt.Println("Площадь круга:", area)
	fmt.Println("Длина окружности:", circumference)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.Radius
}
