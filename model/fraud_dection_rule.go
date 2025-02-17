package model

type FraudDetectionRuleProcessor interface {
	IsFraud(transaction *Transaction) bool
}
