package main

import "fmt"

func main() {
	fmt.Println("Total Sum of the two number is :", calculateSum(20, 70), "Done!")
	// no need to give the format specifiers in Println, it will pick automatically.  println gives the space after the value and hits next line
	// does not support any format specifier.
	fmt.Printf("Total Sum of the two number is :%v ", calculateSum1(20, 30))
	// support the format specifier for printing %v(default format specifier) %d %f.
	// Needs to give the format specifier and add \n for new line and it does not give space after the value
	fmt.Println("Total Sum of the two number is :", calculateSum2(40, 60))

	a, b := calculateSum3(100, 100)
	fmt.Print("Total Sum of the two number is", a, b)
	// Do not support the format specifier for printing.
	// Needs to give the format specifier and add \n for new line and it does not give space after the value
}

// function having 2 parameter way 1
func calculateSum(a, b float64) float64 {
	return a + b
}

// function having 2 parameter and same type way 2
func calculateSum1(a float64, b float64) float64 {
	return a + b
}

// function having 2 parameter and diff type way 2
func calculateSum2(a float64, b int) float64 {
	return a + float64(b)
}

// function having 2 parameter and diff type and multiple return values
func calculateSum3(a float64, b int) (float64, float64) {
	return a + float64(b), a
}
