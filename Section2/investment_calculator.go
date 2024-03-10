package main

import (
	"fmt"
	"math"
)

func main() {
	// Constants
	const inflationRate float64 = 2.5

	// Defined variables
	var investmentAmount float64 = 1000
	var expectedReturnRate float64 = 5.5
	var years int = 10

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Rate of Return: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	var futureValue float64 = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	var futureRealValue float64 = futureValue / math.Pow(1+inflationRate/100, float64(years))
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
