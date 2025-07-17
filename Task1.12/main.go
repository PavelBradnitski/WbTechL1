package main

import "fmt"

func main() {
	var strSlice = []string{"cat", "cat", "dog", "cat", "tree"}
	uniqueStringsSlice := uniqueStrings(strSlice)
	fmt.Println(uniqueStringsSlice)
}

func uniqueStrings(strSlice []string) []string {
	// Создаем map для ускоренного поиска
	uniqueStringsMap := make(map[string]bool)
	for _, v := range strSlice {
		uniqueStringsMap[v] = true
	}
	// Переносим данные из Map в []int
	var uniqueStringsSlice []string
	for key := range uniqueStringsMap {
		uniqueStringsSlice = append(uniqueStringsSlice, key)
	}
	return uniqueStringsSlice
}
