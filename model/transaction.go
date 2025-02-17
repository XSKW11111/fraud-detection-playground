package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Transaction struct {
	ID           uuid.UUID              `json:"id"`
	UserID       uuid.UUID              `json:"user_id"`
	Amount       decimal.Decimal        `json:"amount"`
	MerchantName string                 `json:"merchant_name"`
	Timestamp    *timestamppb.Timestamp `json:"timestamp"`
}
