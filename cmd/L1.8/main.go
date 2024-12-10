package main

import (
	"fmt"
)

// Установить i-й бит в 1
func setBit(num int64, i uint) int64 {
	return num | (1 << i)
}

// Установить i-й бит в 0
func clearBit(num int64, i uint) int64 {
	return num &^ (1 << i)
}

func main() {
	var num int64 = 42 // Пример числа
	var i uint = 3     // Индекс бита, который нужно изменить

	fmt.Printf("Исходное число: %d (%b)\n", num, num)

	// Устанавливаем бит в 1
	num = setBit(num, i)
	fmt.Printf("После установки %d-го бита в 1: %d (%b)\n", i, num, num)

	// Устанавливаем бит в 0
	num = clearBit(num, i)
	fmt.Printf("После установки %d-го бита в 0: %d (%b)\n", i, num, num)
}
