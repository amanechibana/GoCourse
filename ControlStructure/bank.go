package main

import (
	"fmt"
	"os"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile(){
	data, _ := os.ReadFile(accountBalanceFile)
}

func writeBalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile,[]byte(balanceText),0644)
}

func main() {
	var accountBalance = 1000.0 

	for{
		fmt.Println(`Welcome to Pog Bank!
	What do you want to do?
	1. Check Balance
	2. Deposit Money
	3. Withdraw Money
	4. Exit`)

	var choice int
	fmt.Print("Enter: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Your balance is:", accountBalance)
	case 2:
		fmt.Print("Your deposit: ")
		var depositamt float64
		fmt.Scan(&depositamt)

		if depositamt <= 0{
			fmt.Println("Invalid amount. Must be greater than 0!")
			continue 
		}

		accountBalance += depositamt
		fmt.Println("Balance updated! Amount is now: ",accountBalance)
		writeBalanceToFile(accountBalance)
	case 3:
		fmt.Print("Your withdrawal: ")
		var withdrawamt float64
		fmt.Scan(&withdrawamt)

		if withdrawamt <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0!")
			continue
		}

		if withdrawamt > accountBalance{
			fmt.Println("Not enough money!")
			continue 
		}

		accountBalance -= withdrawamt
		fmt.Println("Balance updated! Amount is now: ",accountBalance)
		writeBalanceToFile(accountBalance)
	case 4:
		fmt.Println("Thank you for working with GoBank! Goodbye!")
		return 
	default:
		println("Enter valid number!")
	}

	}
}