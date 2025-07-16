package main

import (
	"fmt"
	"log"
)

func main() {
	var s = []int{1, 2, 3, 4, 5}
	fmt.Println("До удаления:", s)
	s, err := DeleteElem(s, 6)
	if err != nil {
		log.Fatalf("Fail to delete elem: %v", err)
	}
	fmt.Println("После удаления:", s)
}

func DeleteElem(slice []int, index int) ([]int, error) {
	if index < 0 || index > len(slice) {
		return nil, fmt.Errorf("index out of range")
	}
	fmt.Printf("Длина слайса до удаления: %d\n", len(slice))
	copy(slice[index:], slice[index+1:])
	slice = slice[:len(slice)-1]
	fmt.Printf("Длина слайса после удаления: %d\n", len(slice))
	return slice, nil
}
