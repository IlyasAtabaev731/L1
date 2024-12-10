package main

import (
	"fmt"
	"sort"
)

// binarySearch выполняет бинарный поиск в отсортированном срезе
func binarySearch(arr []int, target int) int {
	// Сортируем срез, если он не отсортирован
	sort.Ints(arr)

	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2

		// Если элемент найден, возвращаем его индекс
		if arr[mid] == target {
			return mid
		}
		// Если целевой элемент больше середины, сужаем поиск к правой части
		if arr[mid] < target {
			low = mid + 1
		} else {
			// Иначе сужаем поиск к левой части
			high = mid - 1
		}
	}
	// Если элемент не найден, возвращаем -1
	return -1
}

func main() {
	// Пример данных
	arr := []int{10, 5, 3, 8, 7, 2, 1}
	target := 7

	// Вызов бинарного поиска
	index := binarySearch(arr, target)

	// Вывод результата
	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d\n", target, index)
	} else {
		fmt.Println("Элемент не найден")
	}
}
