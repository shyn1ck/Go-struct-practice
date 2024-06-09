package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func main() {
	book1 := Book{
		Title:  "Война и мир",
		Author: "Лев Толстой",
		Pages:  1225,
	}

	book2 := Book{
		Title:  "Преступление и наказание",
		Author: "Федор Достоевский",
		Pages:  656,
	}

	fmt.Println("Название:", book1.Title)
	fmt.Println("Автор:", book1.Author)
	fmt.Println("Количество страниц:", book1.Pages)
	fmt.Println()

	fmt.Println("Название:", book2.Title)
	fmt.Println("Автор:", book2.Author)
	fmt.Println("Количество страниц:", book2.Pages)
}
