package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	// Увеличиваем счетчик WaitGroup на количество элементов в массиве
	wg.Add(len(numbers))

	for _, num := range numbers {
		// Запускаем горутину для каждого числа
		go func(n int) {
			defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении горутины
			square := n * n
			fmt.Printf("Квадрат числа %d равен %d\n", n, square)
		}(num)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
}
