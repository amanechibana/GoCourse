package prices

type PostTaxPriceJob struct {
	TaxRate float64
	Prices []float64
	PostTaxPrice map[string]float64
}

func NewTaxIncludedPriceJob() *PostTaxPriceJob {
	return &PostTaxPriceJob{
		TaxRate: ,
		Prices: []float64{10,20,30} 
	}
}