package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {

  var investmentAmount float64
  var expectedReturnRate float64
  var years float64

  fmt.Print("Investment Amount: ")
  fmt.Scan(&investmentAmount)

  fmt.Print("Expected return rate: ")
  fmt.Scan(&expectedReturnRate)

  fmt.Print("Number of years: ")
  fmt.Scan(&years)

  futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

  //fmt.Printf("Future value: %.2f\n", futureValue)
  //fmt.Printf("Inflation adapted value: %.2f\n", futureRealValue)

  formattedFV := fmt.Sprintf("Future value: %.2f\n", futureValue)
  formattedRFV := fmt.Sprintf("Inflation adapted value: %.2f\n", futureRealValue)

  fmt.Print(formattedFV, formattedRFV)


}

// (fv float64, rfv float64) = declaring the return variables and their types
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
  fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years) //the := is not necessary as line above creates the fv,rfv variables
  rfv = fv / math.Pow(1 + inflationRate / 100, years)
  return fv, rfv
  //when declaring the variables in function next to type (fv,rfv) then the alternative return could be the one below. 
  //this declaration of variables in the function line, automatically creates them and will be aware of them to return
  //return 
}
