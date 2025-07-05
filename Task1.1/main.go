package main

import "fmt"

// Создаем структуру Human
type Human struct {
	Name string
	Age  int
}

// Реализуем метод getHuman для структуры Human
func (h Human) GetName() string {
	return fmt.Sprintf("Name is: %s, Age is: %d", h.Name, h.Age)
}

// Встраиваем структуру Human
type Action struct {
	Human
}

func main() {
	h := Human{Name: "Pavel", Age: 25}
	a := Action{Human: h}
	// Вызываем метод GetName у структуры Action
	fmt.Println(a.GetName())
}
