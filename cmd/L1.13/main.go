package main

import "fmt"

func main() {
	a := 5
	b := 10

	fmt.Println("До обмена: a =", a, "b =", b)

	// Меняем местами с помощью XOR
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("После обмена: a =", a, "b =", b)
}
