package main

import "fmt"

type Product struct{
	title string
	id int
	price int 
}

func main() {
	var hobbies [3]string = [3]string{"bmt", "guitar", "coding"}

	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	hobbyslice := hobbies[:2]
	hobbyslice2 := hobbies[0:2]
	println(hobbyslice)
	println(hobbyslice2)

	hobbyslice = hobbyslice[2:]
	println(hobbyslice)

	goals := []string{"learn go", "become pro"}
	goals[1] = "olympics 2024"
	goals = append(goals, "become pro")

	fmt.Println(goals)

	food1 := Product{
		"Apple",
		0,
		200,
	}
	food2 := Product{
		"Orange",
		1,
		201,
	}

	food3 := Product{
		"Dragonfruit",
		2,
		202,
	}

	products := []Product{food1,food2}

	fmt.Println(products)

	products = append(products, food3)

	fmt.Println(products)

	prices := []float64{10,11,12,13}

	discountPrices := []float64{1,2,3,4,5}
	prices = append(prices,discountPrices...)

	fmt.Println(prices)
}