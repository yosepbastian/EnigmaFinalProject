package models

type PortFolios struct {
	Id       string  `json:"id"`
	UserID   string  `json:"userid"`
	StockID  int     `json:"stockid"`
	Quantity float64 `json:"quantity"`
}
