package models

type StockSummary struct{

	LongName string
	EpsForward float64
	TrailingAnnualDividendRate  float64
	TrailingAnnualDividendYield float64
	TrailingPE float64
	ForwardPE float64
	PriceToBook float64
	SharesOutstanding int
	MarketCap int64
}