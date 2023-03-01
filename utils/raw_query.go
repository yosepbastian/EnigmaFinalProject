package utils

const (

	SELECT_STOCK_NAME = "SELECT name, price FROM stocks where name=$1"
	SELECT_QUANTITY_STOCK_USER = "SELECT quantity FROM portfolios WHERE user_id=$1 AND stock_id=$2"
	UPDATE_QUANTITY_STOCK_USER = "UPDATE portfolios SET quantity=$1 WHERE user_id=$2 AND stock_id=$3"
	DELETE_STOCK_USER = "DELETE FROM portfolios WHERE user_id=$1 AND stock_id=$2"
	UPDATE_USER_BALANCE = "UPDATE users SET balance=$1 WHERE id=$2"
	UPDATE_QUANTITY_STOCK = "UPDATE stocks SET quantity=$1 WHERE id=$2"
	INSERT_NEW_TRANSACTION = "INSERT into transactions(user_id, stock_id, quantity, price, transaction_type) values ($1,$2,$3,$4,$5)"
	GET_USER_BALANCE = "SELECT balance FROM users WHERE id=$1"
	GET_STOCK_PRICE_BY_ID = "SELECT price from stocks where id=$1"
	GET_STOCK_QUANTITY_BY_ID = "SELECT quantity from stocks where id=$1"

)