package main

import (
	"fmt"
	"time"
)

func main() {
	// Определяем длительность работы программы в секундах
	duration := 10 * time.Second

	// Создаем канал для передачи данных
	ch := make(chan int)

	// Запускаем горутину для отправки данных в канал
	go sendValues(ch, duration)

	// Запускаем горутину для чтения данных из канала
	go readValues(ch, duration)

	// Ждем завершения работы программы
	time.Sleep(duration)
	fmt.Println("Программа завершена")
}

// Функция для отправки значений в канал
func sendValues(ch chan<- int, duration time.Duration) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	value := 0
	for {
		select {
		case <-ticker.C:
			ch <- value
			value++
		case <-time.After(duration):
			return
		}
	}
}

// Функция для чтения значений из канала
func readValues(ch <-chan int, duration time.Duration) {
	for {
		select {
		case value := <-ch:
			fmt.Printf("Получено значение: %d\n", value)
		case <-time.After(duration):
			return
		}
	}
}
