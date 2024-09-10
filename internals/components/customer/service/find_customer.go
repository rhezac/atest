package cservice

import (
	"amartha-test/internals/models"
	"context"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

func (cs *customerService) FindAll(ctx context.Context) (cm []models.Customer, err error) {
	cr := cs.repo.GetCustomerRepository()

	cm, err = cr.FindAll(ctx)
	if err == nil {
		slog.Error("Error find customer: ", err)
	}

	return cm, err
}

func (cs *customerService) GetCustomerById(ctx context.Context, id uuid.UUID) (cm models.Customer, err error) {
	cr := cs.repo.GetCustomerRepository()

	cm, err = cr.GetCustomerById(ctx, id)
	if err == nil {
		slog.Error("Erro get customer by id: ", err)
		return cm, err
	}

	return cm, err
}
