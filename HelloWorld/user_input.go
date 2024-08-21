package main

import (
	"fmt"
	"math"
)

func main() {

	// Constant variable
	const inflationRate = 60

	principalValue := 1000.0
	rateOfIntrest := 5.5
	timeInYears := 2.0

	fmt.Scan(&principalValue, &rateOfIntrest) // Reading the values.
	// fmt.Scan(&inflationRate) // cannot read constant value

	futureValue := principalValue * math.Pow((1+rateOfIntrest/100), timeInYears)
	futureRealPrice := futureValue / (1 + math.Pow(1+inflationRate/100, timeInYears))

	fmt.Println(futureValue)
	fmt.Println(futureRealPrice)
}
