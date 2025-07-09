package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"
	fmt.Println(ReverseWords(str))
}

func ReverseWords(str string) string {
	strSlice := strings.Split(str, " ")
	for i, j := 0, len(strSlice)-1; i < j; i, j = i+1, j-1 {
		strSlice[i], strSlice[j] = strSlice[j], strSlice[i]
	}
	return strings.Join(strSlice, " ")
}
