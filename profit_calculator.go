package main

import (
	"errors"
	"fmt"
	"os"
)

const valuesFile = "values.txt"

func main() {
	var revenue, expense, taxRate float64
	
	//Getting Inputs
	revenue = printAndScan("Enter Revenue: ", revenue)
	revenue, err := validateInput(revenue)

	if err != nil {
		panic(err)
	}

	expense = printAndScan("Enter Expense: ", expense)
	expense, err = validateInput(expense)

	if err != nil {
		panic(err)
	}

	taxRate = printAndScan("Enter Tax Rate: ", taxRate)
	taxRate, err = validateInput(taxRate)

	if err != nil {
		panic(err)
	}
	
	//Calculation
	earningBeforeTax, earningAfterTax, ratio   := calculateValues(revenue, expense, taxRate)

	//Storing Output
	formattedOutput := fmt.Sprintf("Your earnings before tax is: %v\nand your earnings after tax (profit) is: %v\nwhereas the profit to earning ratio is: %.2f", earningBeforeTax, earningAfterTax, ratio)
	storeValue(formattedOutput)

	//Printing Output
	fmt.Printf(formattedOutput)
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

func validateInput(input float64) (float64, error){
	if input <= 0 {
		err := errors.New("invalid input! number should be greater than zero")
		return input, err
	} 
	return input, nil
}

func storeValue(value string){
	os.WriteFile(valuesFile, []byte(value), 0641)
}