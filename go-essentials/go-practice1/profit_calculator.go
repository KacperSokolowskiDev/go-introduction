package main

import (
	"fmt"
)

func main() {

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	
	fmt.Printf("ebt: %.2f, profit: %.2f, ratio: %.2f\n", ebt, profit, ratio)
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64)  {
	ebt := revenue - expenses
	profit := ebt - (ebt * (taxRate/100))
	ratio := ebt / profit
	return ebt, profit, ratio
}
