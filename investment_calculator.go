package main

import (
	"fmt"
	"math"
)

func main() {
  const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64

  fmt.Print("Enter Investment Amount: ")
  fmt.Scan(&investmentAmount)

  fmt.Print("Enter Number of Years: ")
  fmt.Scan(&years)

  fmt.Print("Enter Expected Return Rate: ")
  fmt.Scan(&expectedReturnRate)

  //Formula
  futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years) 
  futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

  fmt.Println(futureValue)
  fmt.Println(futureRealValue)
}