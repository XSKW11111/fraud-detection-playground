package repository

import (
	"database/sql"

	"altas.com/fraud/model"
	"github.com/shopspring/decimal"
)

type MerchantRepository struct {
	db *sql.DB
}

func NewMerchantRepository(db *sql.DB) *MerchantRepository {
	return &MerchantRepository{db: db}
}

func (r *MerchantRepository) initMerchantTable() error {
	_, err := r.db.Exec("CREATE TABLE IF NOT EXISTS merchants (id UUID PRIMARY KEY, name VARCHAR(255), is_high_risk BOOLEAN)")
	return err
}

func (r *MerchantRepository) CreateMerchant(merchant *model.Merchant) error {
	_, err := r.db.Exec("INSERT INTO merchants (id, name, is_high_risk) VALUES (?, ?, ?)", merchant.ID, merchant.Name, merchant.IsHighRisk)
	return err
}

func (r *MerchantRepository) GetMerchant(id string) (*model.Merchant, error) {
	var merchant model.Merchant
	err := r.db.QueryRow("SELECT id, name, is_high_risk FROM merchants WHERE id = ?", id).Scan(&merchant.ID, &merchant.Name, &merchant.IsHighRisk)
	return &merchant, err
}

func (r *MerchantRepository) GetMerchantByName(name string) (*model.Merchant, error) {
	var merchant model.Merchant
	err := r.db.QueryRow("SELECT id, name, is_high_risk FROM merchants WHERE name = ?", name).Scan(&merchant.ID, &merchant.Name, &merchant.IsHighRisk)
	return &merchant, err
}

func (r *MerchantRepository) GetMerchantTransactionCount(id string) (int64, error) {
	var count int64
	err := r.db.QueryRow("SELECT COUNT(*) FROM transactions WHERE merchant_id = ?", id).Scan(&count)
	return count, err
}

func (r *MerchantRepository) GetMerchantTransactionAverageAmount(id string) (decimal.Decimal, error) {
	var amount decimal.Decimal
	err := r.db.QueryRow("SELECT AVG(amount) FROM transactions WHERE merchant_id = ?", id).Scan(&amount)
	return amount, err
}

func (r *MerchantRepository) UpdateMerchantAsHighRisk(merchant *model.Merchant) error {
	_, err := r.db.Exec("UPDATE merchants SET is_high_risk = ? WHERE id = ?", true, merchant.ID)
	return err
}
