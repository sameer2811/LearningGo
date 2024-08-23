package main

import "fmt"

type transformType func(int) int

func main() {

	arr := [5]int{}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Please enter the value for the index %d ", i)
		fmt.Scan(&arr[i])
	}

	resultList := transformNumbers(&arr, operationOddOrEven)
	resultList2 := transformNumbers2(&arr, doubleNumber)
	resultList3 := transformNumbers3(&arr, tripleNumber)

	fmt.Println(resultList, resultList2, resultList3)

}

func transformNumbers(arr *[5]int, operation func(int) int) []int {
	resultNumber := []int{}
	for i := 0; i < len(*arr); i++ {
		resultNumber = append(resultNumber, operation((*arr)[i]))
	}
	return resultNumber
}

func transformNumbers2(arr *[5]int, operation transformType) []int {
	resultNumber := []int{}
	for i := 0; i < len(*arr); i++ {
		resultNumber = append(resultNumber, operation((*arr)[i]))
	}
	return resultNumber
}

func transformNumbers3(arr *[5]int, operation transformType) []int {
	resultNumber := []int{}
	for i := 0; i < len(*arr); i++ {
		resultNumber = append(resultNumber, operation((*arr)[i]))
	}
	return resultNumber
}

func operationOddOrEven(val int) int {
	return val % 2
}

func doubleNumber(value int) int {
	return 2 * value
}

func tripleNumber(value int) int {
	return 3 * value
}
