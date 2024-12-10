package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	data := make(map[int]string)

	wg := sync.WaitGroup{}

	// Конкурентная запись в map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			data[i] = fmt.Sprintf("Value %d", i)
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	// Вывод результатов
	for k, v := range data {
		fmt.Printf("Key: %d, Value: %s\n", k, v)
	}
}
