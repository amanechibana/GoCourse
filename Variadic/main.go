package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	fmt.Println(sumUp(1,10,15,20,25))
	fmt.Println(sumUp(numbers...))
}

func sumUp(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}