package main

import "fmt"

func main (){
	var choice int
	amount   := 100.0

	fmt.Println("Welcome to Go Bank")
	fmt.Println("What do you want to do?")

	for{
    fmt.Println("1. Check Balance")
    fmt.Println("2. Make Deposit")
    fmt.Println("3. Withdraw Amount")
    fmt.Println("4. Exit")
  
    fmt.Print("Your choice: ")
    fmt.Scan(&choice)

		if choice == 1{
			fmt.Print("Your balance is: ", amount)
		} else if choice == 2 {
			var deposit float64
			fmt.Print("Enter the amount you want to deposit: ")
			fmt.Scan(&deposit)
	
			if deposit <= 0 {
				fmt.Println("Invalid Input")
				return
			}
	
			amount  += deposit
			fmt.Print("Deposit successful! Your new is balance is: ", amount)
		} else if  choice == 3 {
			var withdrawalAmount float64
			fmt.Print("Enter the amount you want to withdraw: ")
			fmt.Scan(&withdrawalAmount)
	
			if withdrawalAmount > amount {
				fmt.Println("Error! Your balance is low")
				return
			} 
	
			if withdrawalAmount < 0 {
				fmt.Println("Invalid amount!")
				return
			}
	
			amount -= withdrawalAmount
			fmt.Print("Money withdrawn! New balance: ", amount)
	
		} else {
			fmt.Println("Goodbye!")
      return
		}
	}
}