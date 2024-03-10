package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	getInputFromUser("Revenue: ", &revenue)
	getInputFromUser("Expenses: ", &expenses)
	getInputFromUser("Tax Rate: ", &taxRate)

	earningsBeforeTax, earningsAfterTax, ratio := calculateProfitAndRatio(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.2f\n", earningsAfterTax)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func getInputFromUser(prompt string, floatVariableToStore *float64) {
	fmt.Print(prompt)
	fmt.Scan(floatVariableToStore)
}

func calculateProfitAndRatio(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	var earningsBeforeTax float64 = revenue - expenses
	var earningsAfterTax float64 = earningsBeforeTax * (1 - taxRate/100)
	var ratio float64 = earningsBeforeTax / earningsAfterTax
	return earningsBeforeTax, earningsAfterTax, ratio
}
