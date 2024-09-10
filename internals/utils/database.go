package utils

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"amartha-test/internals/models"
)

func ConnectDatabase(config *Config) (db *gorm.DB, err error) {

	//dsn := "host=localhost user=postgres password=password dbname=go_blog port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return
	}

	fmt.Println("Connected Successfully to the Database")
	return

}

func PrepareDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.Customer{}, &models.Investor{}, &models.LoanProduct{}, &models.LoanProduct{}, &models.LoanAccount{})

	// insert loan product data
	productRate, _ := decimal.NewFromString("10.00")
	loanProduct := &models.LoanProduct{ID: uuid.New(), LoanName: "Normal Loan", Rate: productRate, TenorUnit: "WEEKLY"}
	db.Create(&loanProduct)

	// insert customer data
	uuidD := uuid.New()
	uuidR := uuid.New()
	customers := []models.Customer{
		{ID: uuidD, Name: "Doni", NIK: "3849283047834756", Phone: "0492847562", Email: "doni@example.com", CreateBy: "System", UpdateBy: "System"},
		{ID: uuidR, Name: "Rika", NIK: "1345563047834756", Phone: "0456347562", Email: "rika@example.com", CreateBy: "System", UpdateBy: "System"},
	}
	db.Create(&customers)

	// insert investor data
	uuidID := uuid.New()
	uuidIR := uuid.New()
	balanceD, _ := decimal.NewFromString("50000000.00")
	balanceR, _ := decimal.NewFromString("30000000.00")
	investors := []models.Investor{
		{ID: uuidID, CustomerID: uuidD, Balance: balanceD},
		{ID: uuidIR, CustomerID: uuidR, Balance: balanceR},
	}
	db.Create(&investors)
}
