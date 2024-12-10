package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	sum := 0
	var wg sync.WaitGroup

	// Увеличиваем счетчик WaitGroup на количество элементов в массиве
	wg.Add(len(numbers))

	for _, num := range numbers {
		// Запускаем горутину для каждого числа
		go func(n int) {
			defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении горутины
			sum += n * n
		}(num)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	fmt.Println(sum)
}
