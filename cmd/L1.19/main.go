package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s) // Преобразуем строку в срез рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Меняем элементы местами
	}
	return string(runes) // Преобразуем обратно в строку
}

func main() {
	// Пример строки
	input := "главрыба"
	reversed := reverseString(input)
	fmt.Println("Оригинал:", input)
	fmt.Println("Перевернуто:", reversed)
}
