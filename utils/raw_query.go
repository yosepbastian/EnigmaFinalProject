package utils

const (
	//USERS
	INSERT_USER             = `INSERT INTO users (id, name, email, password, balance) VALUES (:id, :name, :email, :password, :balance)`
	SELECT_USER_ID          = `SELECT * FROM users where id = $1`
	SELECT_USER_BY_PASSWORD = `SELECT * FROM users where password = $1`
	SELECT_USER_BY_EMAIL    = `SELECT * FROM users where email = $1`
	SELECT_EMAIL_FOR_UPDATE = `SELECT * FROM users WHERE email = $1 FOR UPDATE`
	UPDATE_USER             = `UPDATE users SET balance = :balance WHERE id = :id`
	SELECT_USER_NAME        = `SELECT name FROM users where email = $1`

	//validate user login
	VALIDATE_USER_LOGIN = `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1 AND password=$2)`

	//stocks
	SELECT_STOCKS_BY_NAME = `SELECT * FROM stocks WHERE name=$1`
	UPDATE_STOCKS         = `UPDATE stocks SET quantity=:quantity where name= :name`
	GetAll                = `SELECT * FROM stocks`

	//portfolios
	INSERT_PORTFOLIOS           = `INSERT INTO portfolios (id, user_id, stock_id, quantity) VALUES (:id, :userid, :stockid, :quantity)`
	GET_BY_USER_ID_AND_STOCK_ID = `SELECT * FROM portfolios WHERE user_id = $1 AND stock_id = $2`
	UPDATE_PORTFOLIOS           = `UPDATE portfolios SET quantity = :quantity WHERE user_id = :userid AND stock_id = :stockid`
	GET_USER_PORTFOLIOS         = `select b.name, a.quantity, a.stock_id from portfolios as a left join stocks as b on a.stock_id=b.id where a.user_id=$1`

	GET_SELL_STATUS = `select transaction_type from transactions where transaction_type='SELL' AND stock_id=$1 and user_id=$2 limit 1`

	// get last time sell
	GET_LAST_TIME_SELL = `SELECT created_at FROM transactions where user_id=$1 and stock_id=$2 and transaction_type='SELL' order by created_at Desc limit 1`

	// get average price and total quantity
	GET_AVERAGE_PRICE_AND_QUANTITY = `SELECT AVG(price)::numeric(10,2) as average, sum(quantity) as quantity FROM transactions where user_id=$1 and stock_id=$2 and created_at > $3`

	GET_AVERAGE_PRICE_AND_QUANTITY_NO_SELL = `SELECT AVG(price)::numeric(10,2) as average, sum(quantity) as quantity FROM transactions where user_id=$1 and stock_id=$2`

	//transaction

	INSERT_TRANSACTION = `INSERT INTO transactions (id, user_id, stock_id, quantity, price, transaction_type) VALUES(:id, :userid, :stockid, :quantity, :price, :transactiontype)`

	SELECT_STOCK_NAME          = "SELECT name, price FROM stocks where name=$1"
	SELECT_QUANTITY_STOCK_USER = "SELECT quantity FROM portfolios WHERE user_id=$1 AND stock_id=$2"
	UPDATE_QUANTITY_STOCK_USER = "UPDATE portfolios SET quantity=$1 WHERE user_id=$2 AND stock_id=$3"
	DELETE_STOCK_USER          = "DELETE FROM portfolios WHERE user_id=$1 AND stock_id=$2"
	UPDATE_USER_BALANCE        = "UPDATE users SET balance=$1 WHERE id=$2"
	UPDATE_QUANTITY_STOCK      = "UPDATE stocks SET quantity=$1 WHERE id=$2"
	INSERT_NEW_TRANSACTION     = "INSERT into transactions(user_id, stock_id, quantity, price, transaction_type, id) values ($1,$2,$3,$4,$5,$6)"
	GET_USER_BALANCE           = "SELECT balance FROM users WHERE id=$1"
	GET_STOCK_PRICE_BY_ID      = "SELECT price from stocks where id=$1"
	GET_STOCK_QUANTITY_BY_ID   = "SELECT quantity from stocks where id=$1"
)
