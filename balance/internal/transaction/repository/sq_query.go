package repository

const (
	findUsersIDQuery               = `SELECT count(user_id) FROM balance WHERE user_id IN ($1, $2)`
	createTransactionDecreaseQuery = `UPDATE balance SET amount = amount – $1 WHERE user_id = $2 AND amount >= $1`
	createTransactionIncreaseQuery = `UPDATE balance SET amount = amount + $1 WHERE user_id = $2`
	createTransactionHistoryQuery  = `INSERT INTO transactions
	(transaction_id, source, description, sender_id, recipient_id, currency, amount, created_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING transaction_id`
	getTransactionsCountQuery = `SELECT COUNT(*) as total FROM transactions`
	getTransactionsQuery      = `
	SELECT transaction_id, source, description, sender_id, recipient_id, currency, amount, created_at
	FROM transactions WHERE sender_id = $1 OR recipient_id = $1 OFFSET $2 LIMIT $3`
)
