package models

type UserLogin struct{

	Email string `json:"userName"`
	Password string `json:"userPassword"`
}