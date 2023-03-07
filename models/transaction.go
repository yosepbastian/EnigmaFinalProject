package models

type Transaction struct {
	Id              string  `json:"id"`
	UserID          string  `json:"userid"`
	StockID         int     `json:"stockid"`
	Quantity        float64 `json:"quantity"`
	Price           float64 `json:"price"`
	TransactionType string  `json:"transactiontype"`
}
type TransactionAdm struct {
	Id               string  `json:"id"`
	User_ID          string  `json:"userid"`
	Stock_ID         int     `json:"stockid"`
	Quantity         float64 `json:"quantity"`
	Price            float64 `json:"price"`
	Transaction_Type string  `json:"transactiontype"`
	Created_at       string  `json:"created_at"`
}
