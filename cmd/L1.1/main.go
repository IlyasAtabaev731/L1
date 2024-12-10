package main

import "fmt"

// Определяем структуру Human с полями и методами
type Human struct {
	Name string
	Age  int
}

// Метод для структуры Human
func (h Human) Speak(sentence string) {
	fmt.Printf("%s says: %s\n", h.Name, sentence)
}

// Определяем структуру Action, которая встраивает структуру Human
type Action struct {
	Human      // Встраивание структуры Human
	ActionName string
}

// Метод для структуры Action
func (a Action) PerformAction() {
	fmt.Printf("%s is performing action: %s\n", a.Name, a.ActionName)
}

func main() {
	// Создаем экземпляр структуры Action
	action := Action{
		Human: Human{
			Name: "Ilyas",
			Age:  20,
		},
		ActionName: "Programming",
	}

	// Вызываем метод Speak, который принадлежит структуре Human
	action.Speak("Hello world!")

	// Вызываем метод PerformAction, который принадлежит структуре Action
	action.PerformAction()
}
