package infostock

import (
	"kel1-stockbite-projects/models"

	"github.com/musasianturi/finance-go/equity"
	"github.com/musasianturi/finance-go/quote"
)

func GetPrice(stockName string) (float64, error) {

	q, err := quote.Get(stockName)

	if err != nil {
		return 00.0, err
	}

	price := q.Ask

	return price, nil

}

func GetSummary(stockName string) (models.StockSummary, error) {

	var summary models.StockSummary

	q, err := equity.Get(stockName)

	if err != nil {
		return summary, err
	}

	summary = models.StockSummary{
		LongName:                    q.LongName,
		EpsForward:                  q.EpsForward,
		TrailingAnnualDividendRate:  q.TrailingAnnualDividendRate,
		TrailingAnnualDividendYield: q.TrailingAnnualDividendYield,
		TrailingPE:                  q.TrailingPE,
		ForwardPE:                   q.ForwardPE,
		PriceToBook:                 q.PriceToBook,
		SharesOutstanding:           q.SharesOutstanding,
		MarketCap:                   q.MarketCap,
	}

	return summary, nil

}
