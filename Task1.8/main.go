package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {
	n := flag.Uint("bit", 1, "bit of number to change")
	flag.Parse()

	var i int64 = 5
	iNew, err := SetBit(i, *n, 0)
	if err != nil {
		log.Fatalf("error setting bit: %v", err)
	}
	fmt.Printf("Change to 0: %d(%v)\n", iNew, strconv.FormatInt(iNew, 2))
	iNew, err = SetBit(i, *n, 1)
	if err != nil {
		log.Fatalf("error setting bit: %v", err)
	}
	fmt.Printf("Change to 1: %d(%v)\n", iNew, strconv.FormatInt(iNew, 2))
}

func SetBit(num int64, i uint, bitValue int) (int64, error) {
	if bitValue != 0 && bitValue != 1 {
		return 0, errors.New("bitValue must be 0 or 1")
	}

	if bitValue == 1 {
		// Устанавливаем i-й бит в 1, используя побитовое ИЛИ (|)
		mask := int64(1) << i
		num |= mask
	} else {
		// Устанавливаем i-й бит в 0, используя побитовое И (&^)
		mask := int64(1) << i
		num &^= mask
	}

	return num, nil
}
