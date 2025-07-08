package main

import "fmt"

func main() {
	var a = []int{1, -5, -6, 2, -1, 1, 4, 5}
	fmt.Println(quickSort(a))
	fmt.Println(quickSortWithoutNewArrays(a, 0, len(a)-1))
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	quickSortWithoutNewArrays(arr, 0, len(arr)-1)
	return arr
	// pivot := len(arr) / 2
	// var left []int
	// var right []int
	// for i, v := range arr {
	// 	if i == pivot {
	// 		continue
	// 	}
	// 	if arr[pivot] > v {
	// 		left = append(left, v)
	// 	} else {
	// 		right = append(right, v)
	// 	}
	// }
	// left = quickSort(left)
	// right = quickSort(right)
	// var out []int
	// out = append(out, left...)
	// out = append(out, arr[pivot])
	// out = append(out, right...)
	// return out

}

func quickSortWithoutNewArrays(arr []int, low, high int) []int {
	if low < high {
		pivotIndex := (low + high) / 2
		arr[pivotIndex], arr[high] = arr[high], arr[pivotIndex]

		partitionIndex := partition(arr, low, high)

		quickSortWithoutNewArrays(arr, low, partitionIndex-1)
		quickSortWithoutNewArrays(arr, partitionIndex+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
