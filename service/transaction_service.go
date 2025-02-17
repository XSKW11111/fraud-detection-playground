package service

import (
	"database/sql"
	"fmt"

	"altas.com/fraud/model"
	"altas.com/fraud/service/fraud_detection_service"
)

type TransactionService struct {
	fraudDetectionRulesCollection *fraud_detection_service.FraudDetectionRuleProcessorCollection
}

func NewTransactionService(db *sql.DB) *TransactionService {
	return &TransactionService{
		fraudDetectionRulesCollection: fraud_detection_service.NewFraudDetectionRuleProcessorCollection(db),
	}
}

func (s *TransactionService) ProcessTransaction(transaction *model.Transaction) error {
	for _, rule := range s.fraudDetectionRulesCollection.Rules {
		if rule.IsFraud(transaction) {
			return fmt.Errorf("fraud detected")
		}
	}
	return nil
}
