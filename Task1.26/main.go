package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(isUnique("abcd"))
	fmt.Println(isUnique("abCdefAaf"))
	fmt.Println(isUnique("aabcd"))
}

func isUnique(s string) bool {
	runes := []rune(s)
	for i := range runes {
		runes[i] = unicode.ToLower(runes[i])
	}
	s = string(runes)
	isUnique := true
	seen := make(map[rune]bool)

	for _, char := range s {
		if seen[char] {
			isUnique = false
			break
		}
		seen[char] = true
	}

	return isUnique
}
