package repository

import (
	"amartha-test/internals/components/customer"
	"amartha-test/internals/models"
	"context"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type customerRepo dbrepo

var _ customer.CustomerRepository = (*customerRepo)(nil)

func (cr *customerRepo) Insert(ctx context.Context, cm models.Customer) (err error) {
	err = cr.r.dbs.Create(&cm).Error
	if err != nil {
		slog.Error("Failed to insert customer data to db")
		return err
	}

	return
}

func (cr *customerRepo) FindAll(ctx context.Context) (customers []models.Customer, err error) {
	err = cr.r.dbs.Find(&customers).Error
	if err != nil {
		slog.Error("Failed to find customer data from db")
		return customers, err
	}

	return customers, err
}

func (cr *customerRepo) Update(ctx context.Context, cm models.Customer) (cust models.Customer, err error) {
	err = cr.r.dbs.Where("id = ?", cm.ID).Save(cm).Error
	if err != nil {
		slog.Error("Failed to update customer data to db")
		return cm, err
	}

	return cm, err
}

func (cr *customerRepo) GetCustomerById(ctx context.Context, id uuid.UUID) (cust models.Customer, err error) {
	err = cr.r.dbs.Where("id = ?", id).First(&cust).Error
	if err != nil {
		slog.Error("Failed to find customer by id from db")
		return cust, err
	}

	return cust, err
}
