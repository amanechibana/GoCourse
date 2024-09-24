package main

import "fmt"

func main() {
	var revenue,expenses,taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate(In Percent): ")
	fmt.Scan(&taxRate)

	EBT := revenue-expenses
	profit := EBT*(1.00-taxRate/100)
	ratio := EBT/profit

	fmt.Printf("%0.2f, %0.2f, %0.2f",EBT,profit,ratio)
}
