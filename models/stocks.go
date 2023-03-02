package models

type Stocks struct {
	Id       int     `json:"id"`
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price"`
	Quantity float64 `json:"quantity"`
}
