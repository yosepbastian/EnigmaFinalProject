package models

type UserLogin struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
type AdminLogin struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
