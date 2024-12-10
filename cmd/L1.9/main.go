package main

import (
	"fmt"
)

func main() {
	// Определяем два канала
	numbers := make(chan int)
	results := make(chan int)

	// Запуск горутин
	go generateNumbers([]int{1, 2, 3, 4, 5}, numbers)
	go multiplyByTwo(numbers, results)

	// Чтение данных из второго канала
	for result := range results {
		fmt.Println(result)
	}
}

// Функция записи чисел в первый канал
func generateNumbers(arr []int, out chan<- int) {
	for _, num := range arr {
		out <- num
	}
	close(out)
}

// Функция умножения на 2 и записи во второй канал
func multiplyByTwo(in <-chan int, out chan<- int) {
	for num := range in {
		out <- num * 2
	}
	close(out)
}
