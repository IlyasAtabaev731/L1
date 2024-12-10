package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(v[:100]) // Можно работать с частью строки без создания лишней копии
}

func main() {
	someFunc()
}
