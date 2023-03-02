package models

type Users struct {
	Id       string  `json:"id"`
	Name     string  `json:"name" binding:"required"`
	Email    string  `json:"email" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Balance  float64 `json:"balance"`
}

type UsersLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
