package repository

const (
	findUsersIDQuery       = `SELECT count(user_id) FROM balance WHERE user_id IN ($1, $2)`
	createTransactionQuery = ``
)
