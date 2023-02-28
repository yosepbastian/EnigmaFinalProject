package models

type PortFolios struct {
	Id       int    `json:"id"`
	UserID   string `json:"user_id"`
	StockID  int    `json:"stock_id"`
	Quantity int    `json:"quantity"`
}
