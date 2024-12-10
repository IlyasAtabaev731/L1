package main

import (
	"fmt"
)

// Функция быстрой сортировки
func quicksort(arr []int) []int {
	// Если длина массива 1 или меньше, то он уже отсортирован
	if len(arr) <= 1 {
		return arr
	}

	// Выбираем опорный элемент (в данном случае — первый)
	pivot := arr[0]

	// Разбиваем массив на элементы меньше опорного, равные опорному и большие
	var less []int
	var greater []int
	for _, v := range arr[1:] {
		if v <= pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	// Рекурсивно сортируем подмассивы и объединяем с опорным элементом
	return append(append(quicksort(less), pivot), quicksort(greater)...)
}

func main() {
	// Пример массива
	arr := []int{34, 7, 23, 32, 5, 62}
	fmt.Println("До сортировки:", arr)

	// Сортировка массива
	sortedArr := quicksort(arr)
	fmt.Println("После сортировки:", sortedArr)
}
