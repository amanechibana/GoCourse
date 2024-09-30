package main

import (
	"example.com/pricecalc/filemanager"
	"example.com/pricecalc/prices"
	"example.com/pricecalc/cmdmanager"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt",fmt.Sprintf("result_%.0f.json",taxRate*100),)
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxPriceJob(cmdm,taxRate)
		err := priceJob.Process()

		if err != nil{
			fmt.Println(err)
			return
		}

		priceJob = prices.NewTaxPriceJob(fm,taxRate)
		err = priceJob.Process()

		if err != nil{
			fmt.Println(err)
			return
		}
	}
}