package main

import (
	"fmt"

	"example.com/bank/fileops" //path from go.mod module + path to package
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-----")
		//panic(err) // used to crash the program and generate a error message
	}

	fmt.Println("Welcome to Go Bank")
	fmt.Println("Reach us 24/7:", randomdata.PhoneNumber())

	//infinite loop
	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Thank you for choosing our bank")
			fmt.Println("Go Bank wishes you a happy day!")
			return
		}
		
	}

}