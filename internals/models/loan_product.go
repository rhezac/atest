package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type LoanProduct struct {
	ID        uuid.UUID       `json:"id" gorm:"primary_key;UNIQUE"`
	LoanName  string          `json:"loanName" gorm:"type:varchar(100)"`
	Rate      decimal.Decimal `json:"rate" gorm:"type:decimal(5,2)"`
	TenorUnit string          `json:"tenor_unit"`
}
