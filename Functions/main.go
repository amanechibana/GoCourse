package main

import "fmt"

type transformFn func(int) int //unncessary for this code but in longer function types might help clean code 

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers,double)

	fmt.Println(doubled)
	fmt.Println(transformNumbers(&numbers,triple))
}

//func getTransformerFunction() transformFn {
//	return double
//}

func transformNumbers(numbers *[]int, transform transformFn) []int{
	dNumbers := []int{}

	for _,value := range *numbers{
		dNumbers = append(dNumbers,transform(value))
	}

	return dNumbers
}

func double(number int) int{
	return number*2	
}

func triple(number int) int{
	return number*3
}