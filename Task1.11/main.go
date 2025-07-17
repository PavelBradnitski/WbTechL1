package main

import (
	"fmt"
)

func main() {
	var nums1 = []int{1, 2, 3}
	var nums2 = []int{2, 3, 4}
	intersectionSlice := intersection(nums1, nums2)
	fmt.Println(intersectionSlice)
}

func intersection(nums1, nums2 []int) []int {
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
	intersectionSlice := make([]int, 0)
	for key := range intersection {
		intersectionSlice = append(intersectionSlice, key)
	}
	return intersectionSlice
}
