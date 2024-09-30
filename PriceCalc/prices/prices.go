package prices

import (
	"fmt"
	"example.com/pricecalc/conversion"
	"example.com/pricecalc/iomanager"
)

type TaxPriceJob struct {
	IOManager 	 iomanager.IOManager	 	`json:"-"`
	TaxRate      float64 					`json:"tax_rate"`
	Prices       []float64 					`json:"input_price"`
	PostTaxPrice map[string]string 			`json:"tax_included_price"`
}

func (job *TaxPriceJob) LoadData() error{
	lines,err := job.IOManager.ReadLines()

	if err!= nil {
		return err
	}

	prices,err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	job.Prices = prices

	return nil
} 

func (job *TaxPriceJob) Process() error {
	err :=job.LoadData()

	if err != nil{
		return err
	}

	result := make(map[string]string)

	for _, price := range job.Prices {
		result[fmt.Sprintf("%.2f",price)] = fmt.Sprintf("%.2f",(price * (1 + job.TaxRate)))
	}

	job.PostTaxPrice = result
	return job.IOManager.WriteResult(job)
}

func NewTaxPriceJob(iom iomanager.IOManager, taxRate float64) *TaxPriceJob {
	return &TaxPriceJob{
		IOManager: iom,
		TaxRate: taxRate,
		Prices:  []float64{10, 20, 30},
	}
}