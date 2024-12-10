package main

import "fmt"

// Функция для пересечения двух множеств
func intersection(set1, set2 map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})

	// Перебираем элементы первого множества
	for key := range set1 {
		// Если элемент есть во втором множестве, добавляем в результат
		if _, exists := set2[key]; exists {
			result[key] = struct{}{}
		}
	}
	return result
}

func main() {
	// Пример двух множеств
	set1 := map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}}
	set2 := map[int]struct{}{3: {}, 4: {}, 5: {}, 6: {}}

	// Получаем пересечение
	result := intersection(set1, set2)

	// Выводим результат
	fmt.Println("Пересечение двух множеств:")
	for key := range result {
		fmt.Println(key)
	}
}
