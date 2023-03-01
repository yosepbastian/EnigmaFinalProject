package models

import "time"

type Transaction struct {
	Id              string    `json:"id"`
	UserID          string    `json:"userid"`
	StockID         int       `json:"stockid"`
	Quantity        float64   `json:"quantity"`
	Price           float64   `json:"price"`
	TransactionType string    `json:"transactiontype"`
	CreatedAt       time.Time `json:"createdat"`
}
