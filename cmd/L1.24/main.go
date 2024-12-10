package main

import (
	"fmt"
	"math"
)

// Структура для представления точки
type Point struct {
	x, y float64
}

// Конструктор для создания точки
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Метод для нахождения расстояния между двумя точками
func (p *Point) DistanceTo(other *Point) float64 {
	return math.Sqrt(math.Pow(p.x-other.x, 2) + math.Pow(p.y-other.y, 2))
}

func main() {
	// Создаем две точки с помощью конструктора
	point1 := NewPoint(3, 4)
	point2 := NewPoint(7, 1)

	// Находим и выводим расстояние между точками
	distance := point1.DistanceTo(point2)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
