package main

import "fmt"

func main() {
	var revenue, expense, taxRate float64
	
	//Getting Inputs
	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expense: ")
	fmt.Scan(&expense)

	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&taxRate)

	//Calculation
	earningBeforeTax   := revenue - expense
	earningAfterTax    := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / earningAfterTax

	//Printing Output
	fmt.Println(earningBeforeTax)
	fmt.Println(earningAfterTax)
	fmt.Println(ratio)
}