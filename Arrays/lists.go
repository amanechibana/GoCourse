package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99} //dynamically allocated array 
	fmt.Println(prices[0:1])
	
	fmt.Println(append(prices,5.99))
}

//func main() {
//	var productName [4]string = [4]string{"A Book"}
//	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
//	fmt.Println(prices)
//
//	fmt.Println(productName[0])
//	fmt.Println(prices[1:3])
//	fmt.Println(len(prices[1:3]),cap(prices[1:3]))
//
//	// slices are a window, if a slice is created it is a reference to array. 
//	// you can reslice a slice to check more elements to the right 
//	// len is the slices length while cap is possible length including values to the right. 
//}