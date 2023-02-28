package utils

const (
	//USERS
	INSERT_USER = `INSERT INTO users (id, name, email, password, balance) VALUES (:id, :name, :email, :password, :balance)`


	SELECT_STOCKS_BY_NAME = `SELECT * FROM stocks WHERE name=$1`
	UPDATE_STOCKS_BY_ID   = `UPDATE stocks SET name=:name, price=:price, quantity=:quantity WHERE id=:id`
)
