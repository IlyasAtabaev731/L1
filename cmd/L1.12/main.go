package main

import "fmt"

func main() {
	// Исходная последовательность строк
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество с помощью map
	uniqueSet := make(map[string]struct{})

	// Добавляем элементы в множество
	for _, item := range sequence {
		uniqueSet[item] = struct{}{}
	}

	// Выводим множество
	fmt.Println("Множество уникальных элементов:")
	for key := range uniqueSet {
		fmt.Println(key)
	}
}
