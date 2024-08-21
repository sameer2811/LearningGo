package main

import (
	"fmt"
	"math"
)

func main() {

	/*
		Approach 1
		var principalValue int = 1000
		var rateOfIntrest float64 = 5.23
		var timeInYears int = 5

		var futureValue = float64(principalValue) * math.Pow((1+rateOfIntrest/100), float64(timeInYears))
		fmt.Println(futureValue)
	*/

	// Approach 2
	principalValue := 1000.0
	rateOfIntrest := 5.5
	timeInYears := 2.0

	futureValue := principalValue * math.Pow((1+rateOfIntrest/100), timeInYears)
	fmt.Println(futureValue)
}
