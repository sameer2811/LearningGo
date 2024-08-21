package main

import (
	"fmt"
	"math"
)

func main() {
	var principalValue int = 1000
	var rateOfIntrest float64 = 5.23
	var timeInYears int = 5

	var futureValue = float64(principalValue) * math.Pow((1+rateOfIntrest/100), float64(timeInYears))
	fmt.Println(futureValue)
}
