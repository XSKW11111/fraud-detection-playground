package fraud_detection_service

import (
	"database/sql"

	"altas.com/fraud/model"
	"altas.com/fraud/repository"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type FraudDetectionRuleProcessorCollection struct {
	Rules []model.FraudDetectionRuleProcessor
}

func NewFraudDetectionRuleProcessorCollection(db *sql.DB) *FraudDetectionRuleProcessorCollection {
	return &FraudDetectionRuleProcessorCollection{
		Rules: []model.FraudDetectionRuleProcessor{
			&FraudDetectionRuleProcessorOne{
				userAverageTransactions: make(map[uuid.UUID]decimal.Decimal),
				transactionRepo:         repository.NewTransactionRepository(db),
				merchantRepo:            repository.NewMerchantRepository(db),
				userRepo:                repository.NewUserRepository(db),
			},
			&FraudDetectionRuleProcessorTwo{
				transactionRepo: repository.NewTransactionRepository(db),
				merchantRepo:    repository.NewMerchantRepository(db),
				userRepo:        repository.NewUserRepository(db),
			},
			&FraudDetectionRuleProcessorThree{
				userRepo:        repository.NewUserRepository(db),
				merchantRepo:    repository.NewMerchantRepository(db),
				transactionRepo: repository.NewTransactionRepository(db),
			},
			&FraudDetectionProcessorFour{
				transactionRepo: repository.NewTransactionRepository(db),
				merchantRepo:    repository.NewMerchantRepository(db),
			},
			&FraudDetectionProcessorFive{
				transactionRepo: repository.NewTransactionRepository(db),
				merchantRepo:    repository.NewMerchantRepository(db),
			},
		},
	}
}
