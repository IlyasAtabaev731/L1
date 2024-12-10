package main

import (
	"fmt"
	"reflect"
)

func determineType(i interface{}) {
	// Используем типовое утверждение для проверки типа
	switch v := i.(type) {
	case int:
		fmt.Println("Тип переменной: int")
	case string:
		fmt.Println("Тип переменной: string")
	case bool:
		fmt.Println("Тип переменной: bool")
	case chan interface{}:
		fmt.Println("Тип переменной: channel")
	default:
		fmt.Println("Неизвестный тип:", reflect.TypeOf(v))
	}
}

func main() {
	var x interface{}

	// Пример с int
	x = 42
	determineType(x)

	// Пример с string
	x = "Hello"
	determineType(x)

	// Пример с bool
	x = true
	determineType(x)

	// Пример с channel
	x = make(chan int)
	determineType(x)
}
