package investment_controller

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

  fmt.Print("Enter Investment Amount: ")
  fmt.Scan(&investmentAmount)

  fmt.Print("Enter Number of Years: ")
  fmt.Scan(&years)

  fmt.Print("Enter Expected Return Rate: ")
  fmt.Scan(&expectedReturnRate)

  //Formula
  fv, rfv := calculateFutureValues(investmentAmount, expectedReturnRate, years)

  fmt.Println(fv)
  fmt.Println(rfv)
}

func calculateFutureValues (investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
  fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
  rfv = fv / math.Pow(1 + inflationRate / 100, years)
  return
}