package models

import "time"

type Transaction struct {
	Id              int       `json:"id"`
	UserID          int       `json:"user_id"`
	StockID         int       `json:"stock_id"`
	Quantity        int       `json:"quantity"`
	Price           float64   `json:"price"`
	TransactionType string    `json:"transaction_type"`
	CreatedAt       time.Time `json:"created_at"`
}
