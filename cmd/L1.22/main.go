package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализация двух больших чисел
	a := new(big.Int)
	b := new(big.Int)

	// Пример значений > 2^20
	a.SetString("1048577", 10) // a > 2^20
	b.SetString("1234567", 10) // b > 2^20

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Printf("Сложение: %s + %s = %s\n", a.String(), b.String(), sum.String())

	// Вычитание
	diff := new(big.Int).Sub(a, b)
	fmt.Printf("Вычитание: %s - %s = %s\n", a.String(), b.String(), diff.String())

	// Умножение
	prod := new(big.Int).Mul(a, b)
	fmt.Printf("Умножение: %s * %s = %s\n", a.String(), b.String(), prod.String())

	// Деление (с проверкой на деление на ноль)
	if b.Cmp(big.NewInt(0)) != 0 {
		quot := new(big.Int).Div(a, b)
		fmt.Printf("Деление: %s / %s = %s\n", a.String(), b.String(), quot.String())
	} else {
		fmt.Println("Ошибка: деление на ноль!")
	}
}
