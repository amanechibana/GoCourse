package main

import (
	"example.com/pricecalc/filemanager"
	"example.com/pricecalc/prices"
	//"example.com/pricecalc/cmdmanager"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fm := filemanager.New("prices.txt",fmt.Sprintf("result_%.0f.json",taxRate*100),)
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxPriceJob(fm,taxRate)
		go priceJob.Process(doneChans[index],errorChans[index])

		//if err != nil{
		//	fmt.Println(err)
		//	return
		//}

		//priceJob = prices.NewTaxPriceJob(fm,taxRate)
		//err = priceJob.Process()
//
		//if err != nil{
		//	fmt.Println(err)
		//	return
		//}
	}

	for index := range taxRates {
		select{
		case err := <- errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <- doneChans[index]:
			fmt.Println("Done!")
		}
	}
}