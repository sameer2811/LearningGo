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

	// Declared and value assigned automatically
	var inflationRate, price float64 = 30.0, 90 // Declaring multiple values
	fmt.Println(inflationRate, price)

	// Another way of declaring variables without specifying datatype
	newVariablesDeclared, dsfdsfs := 23, 213 // compiler will infer the type accordingly

	fmt.Println(newVariablesDeclared, dsfdsfs)

	// value reassigned
	inflationRate = 20.0

	principalValue := 1000.0
	rateOfIntrest := 5.5
	timeInYears := 2.0

	futureValue := principalValue * math.Pow((1+rateOfIntrest/100), timeInYears)
	futureRealPrice := futureValue / (1 + math.Pow(1+inflationRate/100, timeInYears))

	
	fmt.Println(futureValue)
	fmt.Println(futureRealPrice)
}
