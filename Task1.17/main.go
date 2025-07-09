package main

import (
	"fmt"
	"math"
)

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println(binarySearch(arr, 5))
}

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	foundIndex := -1

	for left <= right {
		mid := int(math.Floor(float64(left+right) / 2))
		if arr[mid] == target {
			foundIndex = mid
			break
		} else if arr[mid] < target {
			left = mid + 1

		} else {
			right = mid + 1
		}
	}
	return foundIndex
}
