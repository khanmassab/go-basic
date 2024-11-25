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

    switch choice{
    case 1:
			fmt.Print("Your balance is: ", amount)
    case 2:
      var deposit float64
			fmt.Print("Enter the amount you want to deposit: ")
			fmt.Scan(&deposit)
	
			if deposit <= 0 {
				fmt.Println("Invalid Input")
				continue
			}
	
			amount  += deposit
			fmt.Print("Deposit successful! Your new is balance is: ", amount)
    case 3:
      var withdrawalAmount float64
			fmt.Print("Enter the amount you want to withdraw: ")
			fmt.Scan(&withdrawalAmount)
	
			if withdrawalAmount > amount {
				fmt.Println("Error! Your balance is low")
				continue
			} 
	
			if withdrawalAmount < 0 {
				fmt.Println("Invalid amount!")
				continue
			}
	
			amount -= withdrawalAmount
			fmt.Print("Money withdrawn! New balance: ", amount)
    default:
      fmt.Println("Goodbye!")
      fmt.Println("Thanks for choosing our bank")
      return
    }
  }
}