package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
//	  => Show error message & exit if invalid input is provided
//    - No negative numbers
//    - Not 0
// 2) Store calculated results into file

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	getInputFromUser("Revenue: ", &revenue)
	var _, err = handleNegativeAndZeroInput(revenue)

	if err != nil {
		fmt.Println(err)
		return
	}

	getInputFromUser("Expenses: ", &expenses)

	_, err = handleNegativeAndZeroInput(expenses)

	if err != nil {
		fmt.Println(err)
		return
	}
	getInputFromUser("Tax Rate: ", &taxRate)

	_, err = handleNegativeAndZeroInput(taxRate)

	if err != nil {
		fmt.Println(err)
		return
	}

	earningsBeforeTax, earningsAfterTax, ratio := calculateProfitAndRatio(revenue, expenses, taxRate)

	writeDataToFile(earningsBeforeTax, earningsAfterTax, ratio)

	fmt.Printf("Earnings Before Tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.2f\n", earningsAfterTax)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func writeDataToFile(earningsBeforeTax float64, earningsAfterTax float64, ratio float64) {
	var commaSeperatedValues = fmt.Sprintf("Earnings Before Tax, Earnings After Tax, Ratio\n%.2f,%.2f,%.2f", earningsBeforeTax, earningsAfterTax, ratio)
	os.WriteFile("profitAndRatio.csv", []byte(commaSeperatedValues), 0644)
}

func handleNegativeAndZeroInput(possibleIncorrectInput float64) (float64, error) {
	if possibleIncorrectInput < 0 {
		return 0, errors.New("program does not support a negative input")
	}

	if possibleIncorrectInput == 0 {
		return 0, errors.New("program does not support a zero input")
	}

	return possibleIncorrectInput, nil
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
