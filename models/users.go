package models

type Users struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password int    `json:"password"`
}
