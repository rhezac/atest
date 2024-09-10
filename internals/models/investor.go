package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Investor struct {
	ID         uuid.UUID       `json:"id" gorm:"primary_key;UNIQUE"`
	CustomerID uuid.UUID       `json:"customer_id"`
	Balance    decimal.Decimal `json:"balance" gorm:"type:decimal(15,2)"`
}
