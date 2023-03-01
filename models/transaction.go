package models


type Transaction struct {
<<<<<<< HEAD
	Id              string    `json:"id"`
	UserID          string    `json:"userid"`
	StockID         int       `json:"stockid"`
	Quantity        float64   `json:"quantity"`
	Price           float64   `json:"price"`
	TransactionType string    `json:"transactiontype"`
	CreatedAt       time.Time `json:"createdat"`
=======
	UserID          string       `json:"user_id"`
	StockID         string       `json:"stock_id"`
	Quantity        int       `json:"quantity"`
	Price           float64   `json:"price"`
	TransactionType string    `json:"transaction_type"`
>>>>>>> 04-Musa
}
