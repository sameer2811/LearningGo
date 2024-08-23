package main

import "fmt"

func main() {
	arrNumber := []int{1, 2, 4, 5}

	for i := 0; i < len(arrNumber); i++ {
		fmt.Printf("Value of the number at index %d is %d \n", i, doubleNumber(arrNumber[i]))
	}
}

func doubleNumber(value int) int {
	return 2 * value
}

func tripleNumber(value int) int {
	return 3 * value
}
