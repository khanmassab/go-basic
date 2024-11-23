package main

import "fmt"

func main() {
	var revenue, expense, taxRate float64
	
	//Getting Inputs
	revenue = printAndScan("Enter Revenue: ", revenue)
	expense = printAndScan("Enter Expense: ", expense)
	taxRate = printAndScan("Enter Tax Rate: ", taxRate)

	//Calculation
	earningBeforeTax, earningAfterTax, ratio   := calculateValues(revenue, expense, taxRate)

	//Printing Output
	fmt.Printf("Your earnings before tax is: %v\nand your earnings after tax (profit) is: %v\nwhereas the profit to earning ratio is: %f", earningBeforeTax, earningAfterTax, ratio)
}

func printAndScan(text string, numberToScan float64) (float64) {
	fmt.Print(text)
	fmt.Scan(&numberToScan)
	return numberToScan
}

func calculateValues (revenue, expense, taxRate float64) (earningBeforeTax float64, earningAfterTax float64, ratio float64) {
	earningBeforeTax   = revenue - expense
	earningAfterTax    = earningBeforeTax * (1 - taxRate/100)
	ratio              = earningBeforeTax / earningAfterTax

	return
}