package service

import (
	"database/sql"
	"fmt"

	"log"

	"altas.com/fraud/model"
	"altas.com/fraud/repository"
	"altas.com/fraud/service/fraud_detection_service"
	"github.com/google/uuid"
)

type TransactionService struct {
	fraudDetectionRulesCollection *fraud_detection_service.FraudDetectionRuleProcessorCollection
	transactionRepo               *repository.TransactionRepository
	merchantRepo                  *repository.MerchantRepository
}

func NewTransactionService(db *sql.DB) *TransactionService {
	return &TransactionService{
		fraudDetectionRulesCollection: fraud_detection_service.NewFraudDetectionRuleProcessorCollection(db),
		transactionRepo:               repository.NewTransactionRepository(db),
		merchantRepo:                  repository.NewMerchantRepository(db),
	}
}

func (s *TransactionService) ProcessTransaction(transaction *model.Transaction) error {
	for _, rule := range s.fraudDetectionRulesCollection.Rules {
		if rule.IsFraud(transaction) {
			return fmt.Errorf("fraud detected")
		}
	}

	err := s.transactionRepo.CreateTransaction(transaction)
	if err != nil {
		log.Println("Error creating transaction", err)
	}

	merchant, err := s.merchantRepo.GetMerchantByName(transaction.MerchantName)
	if err != nil {
		log.Println("Error getting merchant", err)
	}

	if merchant == nil {
		err = s.merchantRepo.CreateMerchant(&model.Merchant{
			ID:   uuid.New(),
			Name: transaction.MerchantName,
		})
		if err != nil {
			log.Println("Error creating merchant", err)
		}
	}

	return nil
}
