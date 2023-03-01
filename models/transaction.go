package models


type Transaction struct {
	Id              string       `json:"id"`
	UserID          string       `json:"user_id"`
	StockID         string       `json:"stock_id"`
	Quantity        int       `json:"quantity"`
	Price           float64   `json:"price"`
	TransactionType string    `json:"transaction_type"`
}
