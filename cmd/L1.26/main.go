package main

import (
	"fmt"
	"strings"
)

// Функция для проверки уникальности символов в строке
func areAllCharactersUnique(s string) bool {
	// Приводим строку к нижнему регистру
	s = strings.ToLower(s)

	// Используем множество для отслеживания уже встреченных символов
	seen := make(map[rune]bool)

	// Проходим по каждому символу в строке
	for _, char := range s {
		// Если символ уже встречался, возвращаем false
		if seen[char] {
			return false
		}
		// Отмечаем символ как встреченный
		seen[char] = true
	}

	// Если все символы уникальны, возвращаем true
	return true
}

func main() {
	// Примеры использования функции
	fmt.Println(areAllCharactersUnique("abcd"))      // true
	fmt.Println(areAllCharactersUnique("abCdefAaf")) // false
	fmt.Println(areAllCharactersUnique("aabcd"))     // false
}
