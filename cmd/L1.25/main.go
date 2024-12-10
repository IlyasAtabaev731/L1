package main

import (
	"fmt"
	"time"
)

// Собственная функция sleep
func mySleep(duration time.Duration) {
	// Создаем таймер
	timer := time.NewTimer(duration)

	// Ожидаем сигнала от таймера
	<-timer.C
}

func main() {
	fmt.Println("Start sleeping for 2 seconds...")
	mySleep(2 * time.Second)
	fmt.Println("Woke up after 2 seconds!")
}
