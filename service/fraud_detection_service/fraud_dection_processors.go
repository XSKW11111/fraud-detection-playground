package fraud_detection_service

import (
	"database/sql"
	"log"

	"time"

	"altas.com/fraud/model"
	"altas.com/fraud/repository"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type FraudDetectionRuleProcessorOne struct {
	userAverageTransactions map[uuid.UUID]decimal.Decimal
	transactionRepo         *repository.TransactionRepository
	userRepo                *repository.UserRepository
	merchantRepo            *repository.MerchantRepository
}

// Rule 1: If the amount is 2x greater than or equal to user average tranactions, it is a fraud transaction
func NewFraudDetectionRuleProcessorOne(db *sql.DB) *FraudDetectionRuleProcessorOne {
	return &FraudDetectionRuleProcessorOne{
		userAverageTransactions: make(map[uuid.UUID]decimal.Decimal),
		transactionRepo:         repository.NewTransactionRepository(db),
		userRepo:                repository.NewUserRepository(db),
		merchantRepo:            repository.NewMerchantRepository(db),
	}
}

func (p *FraudDetectionRuleProcessorOne) IsFraud(transaction *model.Transaction) bool {
	userAverageTransactions, ok := p.userAverageTransactions[transaction.UserID]

	if !ok {
		avg, err := p.userRepo.GetUserAverageTransactionAmount(transaction.UserID)
		if err != nil {
			log.Println("Error getting user average transaction amount", err)
		}
		p.userAverageTransactions[transaction.UserID] = avg
		userAverageTransactions = avg
	}

	if transaction.Amount.GreaterThanOrEqual(userAverageTransactions.Mul(decimal.NewFromInt(2))) {
		return true
	}

	p.transactionRepo.CreateTransaction(transaction)

	merchant, err := p.merchantRepo.GetMerchantByName(transaction.MerchantName)
	if err != nil {
		log.Println("Error getting merchant", err)
	}

	if merchant == nil {
		merchant = &model.Merchant{
			ID:         uuid.New(),
			Name:       transaction.MerchantName,
			IsHighRisk: false,
		}
		p.merchantRepo.CreateMerchant(merchant)
	}

	return false
}

// Rule 2: if the user has more than 5 transactions in 5 minutes, it is a fraud transaction
type FraudDetectionRuleProcessorTwo struct {
	transactionRepo *repository.TransactionRepository
	userRepo        *repository.UserRepository
	merchantRepo    *repository.MerchantRepository
}

func NewFraudDetectionRuleProcessorTwo(db *sql.DB) *FraudDetectionRuleProcessorTwo {
	return &FraudDetectionRuleProcessorTwo{
		transactionRepo: repository.NewTransactionRepository(db),
		userRepo:        repository.NewUserRepository(db),
		merchantRepo:    repository.NewMerchantRepository(db),
	}
}

func (p *FraudDetectionRuleProcessorTwo) IsFraud(transaction *model.Transaction) bool {
	userTransactionCountInTimeRange, err := p.userRepo.GetUserTransactionCountInTimeRange(transaction.UserID, transaction.Timestamp.AsTime().Add(-5*time.Minute), transaction.Timestamp.AsTime())
	if err != nil {
		return false
	}

	if userTransactionCountInTimeRange > 5 {
		return true
	}

	p.transactionRepo.CreateTransaction(transaction)

	merchant, err := p.merchantRepo.GetMerchantByName(transaction.MerchantName)
	if err != nil {
		log.Println("Error getting merchant", err)
	}

	if merchant == nil {
		merchant = &model.Merchant{
			ID:         uuid.New(),
			Name:       transaction.MerchantName,
			IsHighRisk: false,
		}
		p.merchantRepo.CreateMerchant(merchant)
	}
	return false
}

// Rule 3 if user has large first time transaction, it is a fraud transaction
type FraudDetectionRuleProcessorThree struct {
	transactionRepo *repository.TransactionRepository
	userRepo        *repository.UserRepository
	merchantRepo    *repository.MerchantRepository
}

func NewFraudDetectionRuleProcessorThree(db *sql.DB) *FraudDetectionRuleProcessorThree {
	return &FraudDetectionRuleProcessorThree{
		transactionRepo: repository.NewTransactionRepository(db),
		userRepo:        repository.NewUserRepository(db),
		merchantRepo:    repository.NewMerchantRepository(db),
	}
}

func (p *FraudDetectionRuleProcessorThree) IsFraud(transaction *model.Transaction) bool {
	userTransactionCount, err := p.userRepo.GetUserTransactionCount(transaction.UserID)
	if err != nil {
		return false
	}

	if userTransactionCount == 0 && transaction.Amount.GreaterThan(decimal.NewFromInt(100000)) {
		return true
	}

	p.transactionRepo.CreateTransaction(transaction)

	merchant, err := p.merchantRepo.GetMerchantByName(transaction.MerchantName)
	if err != nil {
		log.Println("Error getting merchant", err)
	}

	if merchant == nil {
		merchant = &model.Merchant{
			ID:         uuid.New(),
			Name:       transaction.MerchantName,
			IsHighRisk: false,
		}
		p.merchantRepo.CreateMerchant(merchant)
	}

	return false
}

type FraudDetectionProcessorFour struct {
	merchantRepo    *repository.MerchantRepository
	transactionRepo *repository.TransactionRepository
}

//	Rule 4: if the merchant name is in hight risk list, it is a fraud transaction
//
// if the merchant receive large amount for the first time, it is a fraud transaction
func NewFraudDetectionProcessorFour(db *sql.DB) *FraudDetectionProcessorFour {
	return &FraudDetectionProcessorFour{
		merchantRepo:    repository.NewMerchantRepository(db),
		transactionRepo: repository.NewTransactionRepository(db),
	}
}

func (p *FraudDetectionProcessorFour) IsFraud(transaction *model.Transaction) bool {
	merchant, err := p.merchantRepo.GetMerchant(transaction.MerchantName)
	if err != nil {
		log.Println("Error getting merchant", err)
	}

	if merchant.IsHighRisk {
		return true
	}

	transactionCount, err := p.merchantRepo.GetMerchantTransactionCount(transaction.MerchantName)
	if err != nil {
		log.Println("Error getting merchant transaction count", err)
	}

	if transactionCount == 0 && transaction.Amount.GreaterThan(decimal.NewFromInt(100000)) {
		return true
	}

	p.transactionRepo.CreateTransaction(transaction)

	return false
}

// Rule 5: if the merchant receive large amount that are 2 times of the average amount, it is a fraud transaction
type FraudDetectionProcessorFive struct {
	transactionRepo *repository.TransactionRepository
	merchantRepo    *repository.MerchantRepository
}

func NewFraudDetectionProcessorFive(db *sql.DB) *FraudDetectionProcessorFive {
	return &FraudDetectionProcessorFive{
		transactionRepo: repository.NewTransactionRepository(db),
		merchantRepo:    repository.NewMerchantRepository(db),
	}
}

func (p *FraudDetectionProcessorFive) IsFraud(transaction *model.Transaction) bool {
	merchantTransactionAverageAmount, err := p.merchantRepo.GetMerchantTransactionAverageAmount(transaction.MerchantName)
	if err != nil {
		return false
	}

	if transaction.Amount.GreaterThanOrEqual(merchantTransactionAverageAmount.Mul(decimal.NewFromInt(2))) {
		merchant, err := p.merchantRepo.GetMerchant(transaction.MerchantName)
		if err != nil {
			log.Println("Error getting merchant", err)
		}
		p.merchantRepo.UpdateMerchantAsHighRisk(merchant)
		return true
	}

	p.transactionRepo.CreateTransaction(transaction)

	return false
}
