package models

import (
	"time"

	"github.com/google/uuid"
)

type Document struct {
	ID           uuid.UUID `json:"id" gorm:"primary_key;UNIQUE"`
	Name         string    `json:"name" gorm:"type:varchar(50)"`
	DocUrl       string    `json:"doc_url" gorm:"type:varchar(200)"`
	Description  string    `json:"description" gorm:"type:varchar(200)"`
	DocumentDate time.Time `json:"document_date" gorm:"default:CURRENT_TIMESTAMP"`
	UploadBy     string    `json:"am_name" gorm:"type:varchar(50)"`
}
