package main

import (
	"fmt"

	"github.com/pallinder/go-randomdata"
	"massabatic.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main (){
	var choice int
	amount, err := fileops.GetBalanceFromTheFile(accountBalanceFile)

  if err != nil {
    fmt.Println("error")
    fmt.Println(err)
    fmt.Println("-----------")
    // panic("can't continue sorry")
  }

	fmt.Println("Welcome to Go Bank")
	fmt.Println("Reach us 24/7 at: ", randomdata.PhoneNumber())

  for{
    presentOptions()
    fmt.Print("Your choice: ")
    fmt.Scan(&choice)

    switch choice{
    case 1:
			fmt.Println("Your balance is: ", amount)
    case 2:
      var deposit float64
			fmt.Print("Enter the amount you want to deposit: ")
			fmt.Scan(&deposit)
	
			if deposit <= 0 {
				fmt.Println("Invalid Input")
				continue
			}
	
			amount  += deposit
      fileops.WriteBalanceToFile(amount, accountBalanceFile)
			fmt.Println("Deposit successful! Your new is balance is: ", amount)
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
      fileops.WriteBalanceToFile(amount, accountBalanceFile)
			fmt.Println("Money withdrawn! New balance: ", amount)
    default:
      fmt.Println("Goodbye!")
      fmt.Println("Thanks for choosing our bank")
      return
    }
  }
}