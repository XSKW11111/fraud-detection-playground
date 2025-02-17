package fraud_detection_service

import (
	"database/sql"

	"altas.com/fraud/model"
	"altas.com/fraud/repository"
)

type FraudDetectionRuleProcessorCollection struct {
	Rules []model.FraudDetectionRuleProcessor
}

func NewFraudDetectionRuleProcessorCollection(db *sql.DB) *FraudDetectionRuleProcessorCollection {
	return &FraudDetectionRuleProcessorCollection{
		Rules: []model.FraudDetectionRuleProcessor{
			NewFraudDetectionRuleProcessorOne(db),
			&FraudDetectionRuleProcessorTwo{
				transactionRepo: repository.NewTransactionRepository(db),
				userRepo:        repository.NewUserRepository(db),
			},
			&FraudDetectionRuleProcessorThree{
				userRepo: repository.NewUserRepository(db),
			},
		},
	}
}
