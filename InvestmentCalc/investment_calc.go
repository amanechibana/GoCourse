package main

import (
	"math"
	"fmt"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount,years float64 
	returnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&returnRate)

	fmt.Print("Years to Invest: ")
	fmt.Scan(&years)

	futureVal := investmentAmount * math.Pow(1 + returnRate/100,years)
	realFutureVal := futureVal / math.Pow(1 + inflationRate/100,years)

	//fmt.Println("Future Value:", futureVal)
	//fmt.Println("Future Value(adjusted for inflation):", realFutureVal)
	fmt.Printf("Future Value: %.2f\nFuture Value(adjusted for inflation): %.2f\n", futureVal,realFutureVal)
}
