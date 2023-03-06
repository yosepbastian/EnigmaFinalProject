package models

type PortFolios struct {
	Id       string  `json:"id"`
	UserID   string  `json:"userid"`
	StockID  int     `json:"stockid"`
	Quantity float64 `json:"quantity"`
}

type PortFoliosJoin struct {
	NameStock       string  `json:"name" db:"name"`
	Quantity int `json:"quantity" db:"quantity"`
	StockID  int     `json:"stockId" db:"stock_id"`
}

type AvgQtty struct{
	 Avg float64 `json:"average" db:"average"`
	 Qqty int64 `json:"quantity" db:"quantity"`
	 TimeStamp string `json:"timestamp" db:"created_at"`

}
