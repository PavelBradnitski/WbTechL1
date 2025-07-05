package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) GetName() string {
	return fmt.Sprintf("Name is: %s, Age is: %d", h.Name, h.Age)
}

type Action struct {
	Human
}

func main() {
	h := Human{Name: "Pavel", Age: 25}
	a := Action{Human: h}
	fmt.Println(a.GetName())
}
