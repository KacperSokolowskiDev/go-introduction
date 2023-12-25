package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	//[]byte(balanceText) transforms the string into the byte type required by the method below
	//0644 is a file permission type, permits for owner to read and write and others to only view
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile() (float64, error) {
	// _ means that we know the var exists but we don't want to work with it now
	data, err := os.ReadFile(accountBalanceFile)

	// nil stands for when we don't have anhy useful value, err will be an error if it equals nil
	if err != nil {
		return 1000, errors.New("failed to read balance from file")
	}

	balanceText := string(data) // from []byte to string
	balance, err := strconv.ParseFloat(balanceText, 64) // from string to float64

	if err != nil {
		return 1000, errors.New("failed to parse balance from file")
	}

	return balance, nil //returning no errors
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----")
		//panic(err) // used to crash the program and generate a error message
	}

	fmt.Println("Welcome to Go Bank")

	//infinite loop
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)
		//checkBalance := choice == 1

		//no break necessary in each case in GO
		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount, must be greater than zero")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("New Balance: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Your withdrawal: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid withdraw amount, must be greater than zero")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Invalid withdraw amount, cannot exceed account balance")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("New Balance:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Thank you for choosing our bank")
			fmt.Println("Go Bank wishes you a happy day!")
			return
		}
		
	}

}