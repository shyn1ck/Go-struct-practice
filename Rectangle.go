package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func main() {
	rect := Rectangle{
		Width:  5,
		Height: 10,
	}

	area := rect.Area()

	fmt.Println("Площадь прямоугольника:", area)
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}
