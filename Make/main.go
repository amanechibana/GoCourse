package main 

import "fmt"

type floatMap map[string]float64

func(m floatMap) OutputValue() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2,5) // can go past the allocated space of 5 but then itll have to recreate the array 

	userNames[0] ="Julie"

	userNames = append(userNames,"Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.OutputValue()

	for index, value := range userNames{
		fmt.Println("Index:",index)
		fmt.Println("Value:",value)
	}
	// for range userName{} if you dont care about index or value 

	for key,value := range courseRatings{
		fmt.Println("Key:",key)
		fmt.Println("Value:",value)
	}
}