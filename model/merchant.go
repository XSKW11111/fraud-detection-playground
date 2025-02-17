package model

import "github.com/google/uuid"

type Merchant struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name" gorm:"unique"`
	IsHighRisk bool      `json:"is_high_risk"`
}
