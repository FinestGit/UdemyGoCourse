package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const balanceFile string = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloat64FromFile(balanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
	}

	fmt.Println("Welcome to Go Bank")

	for {
		displayOptions()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("How much do you want to deposit: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloat64ToFile(balanceFile, accountBalance)
		case 3:
			var withdrawalAmount float64
			fmt.Print("How much do you want to withdraw: ")
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Invalid amount. You cannot withdraw more than your current balance.")
				continue
			}

			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
			fileops.WriteFloat64ToFile(balanceFile, accountBalance)
		default:
			if choice == 4 {
				fmt.Println("Goodbye!")
				fmt.Println("Thank you for choosing Go Bank!")
				return
			}
		}
	}
}
