package main

import (
	"fmt"
)

func main() {
	var nums1 = []int{1, 2, 3}
	var nums2 = []int{2, 3, 4}
	intersection := make(map[int]bool)
	// Создаем map для ускоренного поиска
	nums1Map := make(map[int]bool)
	for _, v := range nums1 {
		nums1Map[v] = true
	}
	// добавляем в Map те элементы, которые есть в nums2 и nums1, исключая дубликаты
	for _, v := range nums2 {
		if _, ok := nums1Map[v]; ok {
			intersection[v] = true
		}
	}
	// Переносим данные из Map в []int
	var intersectionSlice []int
	for key := range intersection {
		intersectionSlice = append(intersectionSlice, key)
	}
	fmt.Println(intersectionSlice)
}
