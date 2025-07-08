package main

import (
	"fmt"
	"reflect"
)

func main() {
	DefineType("123")
	DefineType(1)
	DefineType(true)
	ch1 := make(chan int)
	DefineType(ch1)
	ch2 := make(chan<- string)
	DefineType(ch2)
	ch3 := make(<-chan bool)
	DefineType(ch3)
	DefineType(1.1)
}

func DefineType(i interface{}) {
	// switch v := v.(type) {
	// Вместо v.(type) использую reflect, чтобы не проверять все виды каналов
	switch v := reflect.ValueOf(i); v.Kind() {
	case reflect.String:
		fmt.Printf("It's a string %v\n", v)
	case reflect.Int:
		fmt.Printf("It's a int %v\n", v)
	case reflect.Bool:
		fmt.Printf("It's a boolean %v\n", v)
	case reflect.Chan:
		fmt.Printf("It's a chan %v\n", v)
	default:
		fmt.Println("Unknown type")
	}
}
