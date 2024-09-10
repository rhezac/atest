package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type LoanAccount struct {
	ID                  uuid.UUID       `json:"id" gorm:"primary_key;UNIQUE"`
	CustomerID          uuid.UUID       `json:"customer_id"`
	LoanProductID       uuid.UUID       `json:"loan_product_id"`
	Rate                decimal.Decimal `json:"rate" gorm:"type:decimal(5,2)"`
	ValidateDate        time.Time       `json:"validate_date"`
	ValidateBy          string          `json:"validate_by" gorm:"type:varchar(50)"`
	ApproveDate         time.Time       `json:"approve_date"`
	AgreementNo         string          `json:"agreement_no" gorm:"varchar(50)"`
	AgreementDate       time.Time       `json:"agreement_date"`
	AgreementDocID      uuid.UUID       `json:"agreement_doc_id"`
	DisbursementDate    time.Time       `json:"disbursement_date"`
	DisbursementOfficer string          `json:"disbursement_officer"`
	Principal           decimal.Decimal `json:"principal_amount" gorm:"type:decimal(15,2)"`
	TotalAmount         decimal.Decimal `json:"total_amount" gorm:"type:decimal(15,2)"`
	OutstandingAmount   decimal.Decimal `json:"outstanding_amount" gorm:"type:decimal(15,2)"`
	ROI                 decimal.Decimal `json:"roi" gorm:"type:decimal(5,2)"`
	TenorValue          int             `json:"tenor_value"`
	Status              string          `json:"loan_status" gorm:"varchar(15)"`
	LoanState           string          `json:"loan_state" gorm:"varchar(15)"`
	Kind                string          `json:"kind" gorm:"json"`
	Investors           []LoanInvested  `json:"investors" gorm:"json"`
	CreateBy            string          `json:"create_by" gorm:"varchar(50)"`
	UpdateBy            string          `json:"update_by" gorm:"varchar(50)"`
	CreateDate          time.Time       `json:"create_date" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateDate          time.Time       `json:"update_date" gorm:"default:CURRENT_TIMESTAMP"`
	Documents           Document        `json:"documents" gorm:"json"`
}

type LoanInvested struct {
	ID           uuid.UUID       `json:"id"`
	InvestorID   uuid.UUID       `json:"investor_id"`
	LoanID       uuid.UUID       `json:"loan_id"`
	InvestAmount decimal.Decimal `json:"invest_amount"`
	CreateAt     time.Time       `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	CreateBy     string          `json:"-" gorm:"varchar(50)"`
}
