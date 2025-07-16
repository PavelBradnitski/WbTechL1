package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Start: %v\n", time.Now())
	Sleep(time.Second * 5)
	fmt.Printf("Finish: %v\n", time.Now())
}

func Sleep(duration time.Duration) {
	t := time.NewTimer(duration)
	<-t.C
}
