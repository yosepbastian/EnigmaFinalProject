package models

type Users struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password int     `json:"password"`
	Balance  float64 `json:"balance"`
}
