package repository

import (
	"database/sql"
	"time"

	"altas.com/fraud/model"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) initUserTable() error {
	_, err := r.db.Exec("CREATE TABLE IF NOT EXISTS users (id UUID PRIMARY KEY, name VARCHAR(255))")
	return err
}

func (r *UserRepository) GetUser(id uuid.UUID) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id FROM users WHERE id = $1", id).Scan(&user.ID)
	return &user, err
}

func (r *UserRepository) GetUserAverageTransactionAmount(userID uuid.UUID) (decimal.Decimal, error) {
	var avg sql.NullString
	err := r.db.QueryRow("SELECT AVG(amount) FROM transactions WHERE user_id = $1", userID).Scan(&avg)
	if err != nil {
		return decimal.Zero, err
	}

	if !avg.Valid {
		return decimal.Zero, nil
	}

	return decimal.NewFromString(avg.String)
}

func (r *UserRepository) GetUserTransactionCount(userID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM transactions WHERE user_id = $1", userID).Scan(&count)
	return count, err
}

func (r *UserRepository) GetUserTransactionCountInTimeRange(userID uuid.UUID, startTime time.Time, endTime time.Time) (int64, error) {
	var count int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM transactions WHERE user_id = $1 AND timestamp BETWEEN $2 AND $3", userID, startTime, endTime).Scan(&count)
	return count, err
}
