package models

type OrderRequest struct {

	UserID    string `json:"userId" binding:"required"`
	Email     string `json:"email" binding:"required"`
	StockName string `json:"stockName" binding:"required"`
	Quantity  float64    `json:"quantity" binding:"required"`
	Price     float64    `json:"price" binding:"required"`
	
}
