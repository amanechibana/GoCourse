package main

import (
	"fmt"
	"errors"
	"os"
)

func main() {
	revenue,err1 := getUserInput("Revenue: ")

	//if err != nil{
	//	fmt.Println(nil)
	//	//return
	//	panic(err)
	//}

	expenses,err2 := getUserInput("Expenses: ")

	//if err != nil{
	//	fmt.Println(nil)
	//	//return
	//	panic(err)
	//}

	taxRate,err3:= getUserInput("Tax Rate: ")

	//if err != nil{
	//	fmt.Println(nil)
	//	//return
	//	panic(err)
	//}

	if err1 != nil || err2 != nil || err3 != nil { 
		fmt.Println(err1)
		panic(err1)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
	saveInfo(ebt,profit,ratio)
}

func saveInfo(EBT, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProft: %.1f\n Ratio:%.1f\n",EBT,profit,ratio)
	os.WriteFile("results.txt",[]byte(results),0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64,error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0{
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput,nil
}
