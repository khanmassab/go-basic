package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func writeBalanceToFile(balance float64){
  balanceText := fmt.Sprint(balance)
  os.WriteFile(accountBalanceFile,[]byte(balanceText), 0644)
}

func getBalanceFromTheFile() (float64, error){
  data, err := os.ReadFile(accountBalanceFile)

  if err != nil {
    return 1000, errors.New("error finding the balance file")
  }
  balanceText := string(data)
  balance, err := strconv.ParseFloat(balanceText, 64)

  if err != nil {
    return 1000, errors.New("error reading the balance file")
  }
  return balance, nil
}

func main (){
	var choice int
	amount, err := getBalanceFromTheFile()

  if err != nil {
    fmt.Println("error")
    fmt.Println(err)
    fmt.Println("-----------")
    // panic("can't continue sorry")
  }

	fmt.Println("Welcome to Go Bank")

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
      writeBalanceToFile(amount)
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
      writeBalanceToFile(amount)
			fmt.Println("Money withdrawn! New balance: ", amount)
    default:
      fmt.Println("Goodbye!")
      fmt.Println("Thanks for choosing our bank")
      return
    }
  }
}