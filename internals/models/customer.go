package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;UNIQUE"`
	Name       string    `json:"name" gorm:"varchar(50)"`
	NIK        string    `json:"nik" gorm:"varchar(20)"`
	Phone      string    `json:"phone" gorm:"varchar(20)"`
	Email      string    `json:"email" gorm:"varchar(50)"`
	CreateBy   string    `json:"create_by" gorm:"varchar(50)"`
	UpdateBy   string    `json:"update_by" gorm:"varchar(50)"`
	CreateDate time.Time `json:"create_date" gorm:"default:CURRENT_TIMESTAMP"`
	UpdateDate time.Time `json:"update_date" gorm:"default:CURRENT_TIMESTAMP"`
}
