package models

type Portfolio struct {
	Id       int `json:"id"`
	UserID   int `json:"user_id"`
	StockID  int `json:"stock_id"`
	Quantity int `json:"quantity"`
}