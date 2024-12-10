package main

import "fmt"

func main() {
	// Исходный слайс
	slice := []int{1, 2, 3, 4, 5}

	// Индекс элемента, который нужно удалить
	i := 2

	// Удаление элемента с индексом i
	slice = append(slice[:i], slice[i+1:]...)

	// Вывод результата
	fmt.Println(slice) // [1 2 4 5]
}
