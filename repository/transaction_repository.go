package repository

import (
	"database/sql"

	"altas.com/fraud/model"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) initTransactionTable() error {
	_, err := r.db.Exec("CREATE TABLE IF NOT EXISTS transactions (id UUID PRIMARY KEY, amount DECIMAL(10, 2), user_id UUID, merchant_name VARCHAR(255), timestamp TIMESTAMP)")
	return err
}

func (r *TransactionRepository) CreateTransaction(transaction *model.Transaction) error {
	_, err := r.db.Exec("INSERT INTO transactions (id, amount, user_id, merchant_name, timestamp) VALUES ($1, $2, $3, $4, $5)", transaction.ID, transaction.Amount, transaction.UserID, transaction.MerchantName, transaction.Timestamp.AsTime().UTC())
	return err
}

func (r *TransactionRepository) GetTransaction(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	err := r.db.QueryRow("SELECT id, amount, user_id FROM transactions WHERE id = $1", id).Scan(&transaction.ID, &transaction.Amount, &transaction.UserID)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
