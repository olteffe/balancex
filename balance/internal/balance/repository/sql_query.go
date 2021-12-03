package repository

const (
	createBalanceQuery = `INSERT INTO balance (user_id, currency, amount, created_at) 
	VALUES ($1, $2, $3, $4) RETURNING user_id`
	getBalanceQuery = `SELECT user_id, currency, amount, updated_at FROM balance WHERE user_id=$1`
)
