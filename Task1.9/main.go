package main

import (
	"fmt"
)

func sendNumsToChan(numSlice []int, out chan<- int) {
	defer close(out)
	for _, v := range numSlice {
		out <- v
	}
}

func multiply(in <-chan int, out chan<- int) {
	defer close(out)
	for val := range in {
		out <- val * 2
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	numSlice := []int{1, 3, 4, 5, 6, 7, 8}
	go sendNumsToChan(numSlice, ch1)
	go multiply(ch1, ch2)

	for val := range ch2 {
		fmt.Println("Received from SendDoubleValue output:", val)
	}
	fmt.Println("Finished")
}
